package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"oldpackage_go_shop/hello_grpc/proto"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {

	return &proto.HelloReply{
		Reply: "hello, " + request.Name,
	}, nil
}

func main() {

	g := grpc.NewServer()

	proto.RegisterGreeterServer(g, &Server{})

	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("failed to listen: ", err)
	}

	g.Serve(lis)

}
