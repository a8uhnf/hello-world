package main

import (
	"context"
	"fmt"
	"os"

	"github.com/a8uhnf/hello-world/grpc-streaming/api"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting client")
	address := os.Getenv("GRPC_SERVER")
	if address == "" {
		panic("empty server address")
	}

	ctx := context.Background()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	echoCli := api.NewEchoServiceClient(conn)

	resp, err := echoCli.Echo(ctx, &api.EchoRequest{})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
