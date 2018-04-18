package main

import (
	"context"
	"fmt"
	"time"

	"github.com/a8uhnf/hello-world/grpc_ext/api"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World!!!")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(":65000", opts...)
	if err != nil {
		panic(nil)
	}
	defer conn.Close()
	client := api.NewHelloHanifaClient(conn)
	c := time.Tick(time.Second * 2)
	for now := range c {
		fmt.Println(now)
		resp, err := client.HelloHanifa(context.Background(), &api.HelloReq{Name: "hanifa"})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}
}
