package main

import (
	"fmt"
	"net"

	"github.com/a8uhnf/hello-world/grpc_ext/api"
	grpc "google.golang.org/grpc"
)

func main() {
	address := fmt.Sprintf(":8080")
	fmt.Println("Localhost...", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	s := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterHelloHanifaServer(grpcServer, &s)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
