package api

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	env := os.Getenv("HELLO_WORLD")
	fmt.Println("---------------", env)
	n, err := strconv.Atoi(env)
	if err != nil {
		return nil, err
	}
	fmt.Println("---------------", n)

	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: env}, nil
}
