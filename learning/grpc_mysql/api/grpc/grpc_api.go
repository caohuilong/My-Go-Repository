package grpc

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"grpc_mysql/api/mysql"
	"log"
	"net"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	pb "grpc_mysql/api/proto"
)

// config是server的配置
type config struct {
	// gRPC服务器启动参数部分
	GRPCPort string //gRPC服务i去监听的端口

	// 数据库配置部分
	DatastoreDBHost     string //数据库地址
	DatastoreDBUser     string //连接数据库的用户名
	DatastoreDBPassword string //连接数据库的密码
	DatastoreDBSchema   string //数据库的名称
}

// 配置参数并调用runServer启动gRPC服务器和HTTP网关
func Run() error {
	ctx := context.Background()

	//获取配置
	var cfg config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.DatastoreDBUser, cfg.DatastoreDBPassword, cfg.DatastoreDBHost, cfg.DatastoreDBSchema)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	mysqlAPI := mysql.NewStaffServiceServer(db)

	return runServer(ctx, mysqlAPI, cfg.GRPCPort)
}

// RunServer运行gRPC服务以提供对外服务接口
func runServer(ctx context.Context, mysqlAPI pb.StaffServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterStaffServiceServer(server, mysqlAPI)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// 启动gRPC服务器
	log.Println("starting gRPC server...")
	if err := server.Serve(listen); err != nil {
		log.Fatal("starting gRPC server failed...")
		return err
	}

	return nil
}
