package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	// client, _ := rpc.Dial("tcp", "localhost:1234")

	// var reply *string = new(string)
	var reply string

	// 前面的dial 拨号，是Gob格式
	conn, _ := net.Dial("tcp", "localhost:1234")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	err := client.Call("HelloService.Hello", "lee", &reply)

	if err != nil {
		fmt.Println("client error: ", err)
		return
	}

	fmt.Println("reply: ", reply)

}
