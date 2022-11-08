package main

import (
	"fmt"
	"grpc_mysql/api/grpc"
	"os"
)

func main() {
	if err := grpc.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
