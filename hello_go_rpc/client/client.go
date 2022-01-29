package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	// client, _ := rpc.Dial("tcp", "localhost:1234")
	client, _ := rpc.Dial("tcp", "106.53.72.208:1234")

	// var reply *string = new(string)
	var reply string
	err := client.Call("HelloService.Hello", "lee, i am from your server", &reply)

	if err != nil {
		fmt.Println("client error: ", err)
		return
	}

	fmt.Println("reply: ", reply)

	
}
