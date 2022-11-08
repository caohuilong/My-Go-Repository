package main

import (
	"fmt"
	"http_grpc/api/http_api"
	"os"
)

func main() {
	if err := http_api.RunHTTP(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
