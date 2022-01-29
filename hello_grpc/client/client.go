package main

import (
	"context"
	"fmt"
	"oldpackage_go_shop/hello_grpc/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	//grpc.WithInsecure)
	defer conn.Close()

	if err != nil {
		fmt.Println("conn failed: ", err)
	}

	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "lee",
	})

	if err != nil {
		fmt.Println("call error: ", err)
	}

	fmt.Println(r.Reply)

}
