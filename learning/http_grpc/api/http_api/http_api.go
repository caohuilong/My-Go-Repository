package http_api

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"http_grpc/api/grpc_api"
	"http_grpc/api/mysql_api"
	pb "http_grpc/api/proto"
)

// Config是Server的配置
type Config struct {
	// gRPC服务器启动参数部分
	GRPCPort string //GRPCPort是gRPC服务器监听的TCP端口

	// HTTP/REST网关启动参数部分
	HTTPPort string //HTTPPort是通过HTTP/REST网关监听的TCP端口

	// 数据库配置部分
	DatastoreDBHost     string //DatestoreDBHost是数据库的地址
	DatastoreDBUser     string //DatastoreDBUser是用于连接数据库的用户名
	DatastoreDBPassword string //DatastoreDBPassword是用于连接数据库的密码
	DatastoreDBSchema   string //DatastoreDBSchema是数据库的名称
}

// Run运行gRPC服务器和HTTP网关
func RunHTTP() error {
	ctx := context.Background()

	// 获取配置
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", cfg.DatastoreDBUser, cfg.DatastoreDBPassword, cfg.DatastoreDBHost, cfg.DatastoreDBSchema)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database:%v", err)
	}
	defer db.Close()

	mysqlAPI := mysql_api.NewStaffServiceServer(db)

	// 运行HTTP网关
	go func() {
		_ = runHTTPServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc_api.RunGRPCServer(ctx, mysqlAPI, cfg.GRPCPort)
}

// runServer运行HTTP/REST网关
func runHTTPServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterStaffServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v\n", err)
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// 优雅关闭
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			<-ctx.Done()
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP gateway...")
	return srv.ListenAndServe()
}
