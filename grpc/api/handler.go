package api

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	env := os.Getenv("HELLO_WORLD")
	fmt.Println("---------------", env)

	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: env}, nil
}
