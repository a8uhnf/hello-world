package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/a8uhnf/hello-world/grpc-streaming/api"
	"google.golang.org/grpc"
)

type EchoServiceHandler struct {
	api.UnimplementedEchoServiceServer
}

var mp map[int64]chan string

var cnt = 0

var userId int64 = 0

func (e *EchoServiceHandler) Echo(ctx context.Context, request *api.EchoRequest) (*api.EchoResponse, error) {

	if _, ok := mp[request.UserId]; !ok {
		return nil, fmt.Errorf("user_id not valid")
	}
	ch := mp[request.UserId]
	ch <- request.Message
	return &api.EchoResponse{
		Message:      "sender successful",
		// MessageCount: int32(cnt),
	}, nil
}

func (e *EchoServiceHandler) EchoAbort(ctx context.Context, request *api.EchoRequest) (*api.EchoResponse, error) {
	cnt = 0

	return &api.EchoResponse{
		Message:      "aborted",
		MessageCount: 0,
	}, nil
}

func (e *EchoServiceHandler) NoOp(ctx context.Context, empty *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}

func (e *EchoServiceHandler) ServerStreamingEcho(request *api.ServerStreamingEchoRequest, server api.EchoService_ServerStreamingEchoServer) error {
	log.Printf("fetch response for id : %s", request.Message)
	ch := make(chan string, 100)
	id := userId
	userId++
	mp = make(map[int64]chan string)

	mp[id] = ch

	log.Printf("user_id: %d", id)

	cnt++

	resp := api.ServerStreamingEchoResponse{
		Message: "registered",
		UserId:  id,
	}
	if err := server.Send(&resp); err != nil {
		log.Printf("send error %v", err)
	}

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
			resp := api.ServerStreamingEchoResponse{
				Message: v,
				UserId:  id,
			}
			if err := server.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("send response to : %d", id)
		}
	}

	// use wait group to allow process to be concurrent
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(count int64) {
	// 		defer wg.Done()
	//
	// 		//time sleep to simulate server process time
	// 		time.Sleep(time.Duration(count) * time.Second)
	// 		resp := api.ServerStreamingEchoResponse{
	// 			Message: "hello world!!!!",
	// 		}
	// 		if err := server.Send(&resp); err != nil {
	// 			log.Printf("send error %v", err)
	// 		}
	// 		log.Printf("finishing request number : %d", count)
	// 	}(int64(i))
	// }

	// wg.Wait()
	log.Printf("finished for user_id %d", id)
	return nil

}

func (e *EchoServiceHandler) ServerStreamingEchoAbort(request *api.ServerStreamingEchoRequest, server api.EchoService_ServerStreamingEchoAbortServer) error {
	panic("implement me")
}

func (e *EchoServiceHandler) ClientStreamingEcho(server api.EchoService_ClientStreamingEchoServer) error {
	panic("implement me")
}

func (e *EchoServiceHandler) FullDuplexEcho(server api.EchoService_FullDuplexEchoServer) error {
	panic("implement me")
}

func (e *EchoServiceHandler) HalfDuplexEcho(server api.EchoService_HalfDuplexEchoServer) error {
	panic("implement me")
}

func main() {
	fmt.Println("starting grpc server")
	address := os.Getenv("GRPC_ADDRESS")
	if address == "" {
		panic("empty grpc address env variable")
	}

	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	// create a server instance
	s := &EchoServiceHandler{}

	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	api.RegisterEchoServiceServer(grpcServer, s)
	// start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
