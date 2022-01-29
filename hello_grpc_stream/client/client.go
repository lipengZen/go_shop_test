package main

import (
	"context"
	"fmt"
	"math/rand"
	"oldpackage_go_shop/hello_grpc_stream/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	//grpc.WithInsecure)

	if err != nil {
		fmt.Println("conn failed: ", err)
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	/* 服务端流模式
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "哈哈"})

	for {
		a, err := res.Recv() // 如果懂socket编程的话，就会明白send recv
		if err != nil {
			fmt.Println("error : ", err)
			return
		}
		fmt.Println("a: ", a)

	}
	*/

	/* 客户端流模式
	req, _ := c.PutStream(context.Background())
	i := 0
	for {
		str := strconv.Itoa(i)
		req.Send(&proto.StreamReqData{Data: "哈哈" + str})
		i++
		time.Sleep(1 * time.Second)
		if i > 20 {
			break
		}
	}
	*/

	cli, _ := c.AllStream(context.Background())

	rand.Seed(time.Now().Unix())
	x := rand.Intn(100)

	_ = cli.Send(&proto.StreamReqData{Data: "hello, "})

	for x < 95 {
		fmt.Println("x: ", x)
		a, err := cli.Recv()
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		fmt.Println("qian: ", a.Data)

		_ = cli.Send(&proto.StreamReqData{Data: "哈哈"})

		x = rand.Intn(100)
		time.Sleep(time.Second)
	}

}
