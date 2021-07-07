package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/a8uhnf/hello-world/grpc-streaming/api"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting client")
	address := os.Getenv("GRPC_ADDRESS")
	if address == "" {
		panic("empty server address")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// create stream
	client := api.NewEchoServiceClient(conn)
	in := &api.ServerStreamingEchoRequest{Message: "hello"}
	stream, err := client.ServerStreamingEcho(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s || user_id: %d", resp.Message, resp.UserId)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")

}
