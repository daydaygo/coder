package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	hellopb "basic/grpc/proto/hello"
)

const Address = ":50052"

type HelloService struct{}

func (h HelloService) SayHello(ctx context.Context, req *hellopb.HelloReq) (*hellopb.HelloResp, error) {
	panic("implement me")
}

func main() {
	l, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hellopb.RegisterHelloServer(s, HelloService{})
	grpclog.Infoln("listen on: " + Address)
	s.Serve(l)
}
