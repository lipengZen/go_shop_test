package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"

	"oldpackage_go_shop/grpclb_test/proto"
)

func main() {
	conn, err := grpc.Dial(
		"consul://127.0.0.1:8500/user_srv?wait=14s&tag=lee", // tag不能写错
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userSrvClient := proto.NewUserClient(conn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 2,
	})

	fmt.Println(rsp, err)

}
