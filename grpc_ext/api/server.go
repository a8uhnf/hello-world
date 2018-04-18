package api

import (
	"fmt"

	"golang.org/x/net/context"
)

type Server struct{}

var count int

func (s Server) HelloHanifa(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	count++
	fmt.Println("-----------------------------------")
	fmt.Println(req.Name)
	return &HelloResp{Count: int32(count)}, nil
}
