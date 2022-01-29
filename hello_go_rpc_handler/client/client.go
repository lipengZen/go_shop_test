package main

import (
	"fmt"

	client_proxy0 "oldpackage_go_shop/hello_go_rpc_handler/client_proxy"
)

func main() {

	// client, _ := rpc.Dial("tcp", "localhost:1234")

	var reply *string = new(string)
	// var reply string

	// // 前面的dial 拨号，是Gob格式
	// conn, _ := net.Dial("tcp", "localhost:1234")
	// client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	client := client_proxy0.NewHelloServiceStub("tcp", "localhost:1234")

	err := client.Hello("from stub", reply)
	// err := client.Call(handler.HelloServiceName+".Hello", "lee", reply)

	if err != nil {
		fmt.Println("client error: ", err)
		return
	}

	fmt.Println("reply: ", *reply)

}
