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

var cnt = 0

func (e *EchoServiceHandler) Echo(ctx context.Context, request *api.EchoRequest) (*api.EchoResponse, error) {
	cnt++
	return &api.EchoResponse{
		Message:      request.Message,
		MessageCount: int32(cnt),
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
	panic("implement me")
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
