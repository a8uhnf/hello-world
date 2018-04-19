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
	// target := os.Getenv("ECHO_SERVICE_TARGET")
	// port := os.Getenv("ECHO_SERVICE_PORT")
	target := "35.199.185.138:8080"
	fmt.Println(target)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s", target), opts...)
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
