package main

import (
	"context"
	"flag"
	pb "grpc_mysql/api/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// 获取配置
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// 建立与服务器的连接
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStaffServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 调用Create
	req1 := pb.CreateRequest{
		Staffinfo: &pb.StaffInfo{
			Name:        "guofenglin",
			Phonenumber: "17724554775",
			Position:    "Front End",
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	// 调用Read
	req2 := pb.ReadRequest{
		Name: "caohuilong",
	}
	res2, err := c.Read(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res2)

	// 调用Update
	req3 := pb.UpdateRequest{
		Staffinfo: &pb.StaffInfo{
			Id:          res2.Staffinfo.Id,
			Name:        res2.Staffinfo.Name,
			Phonenumber: "17328681706",
			Position:    "Algorithm Engineer",
		},
	}
	res3, err := c.Update(ctx, &req3)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res3)

	// 调用ReadAll
	req4 := pb.ReadAllRequest{}
	res4, err := c.ReadAll(ctx, &req4)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res4)

	// 调用Delete
	req5 := pb.DeleteRequest{
		Name: "caohuilong",
	}
	res5, err := c.Delete(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res5)
}
