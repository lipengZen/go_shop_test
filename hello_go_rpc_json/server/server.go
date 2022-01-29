package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{

}

func (s *HelloService) Hello( request string, reply *string) error{
	*reply = "hello, " + request
	return nil
}


func main() {
	// 1. 实例化一个server
	listener, err := net.Listen("tcp",":1234")

	if err!=nil {
		fmt.Println("error: " , err)
		return
	}

	// 2. 注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{} )

	// 3. 启动服务
	conn, _ := listener.Accept()
	// rpc.ServeConn(conn)

	// 改成 新的编解码方式，
	// 很简单：使用了 json
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

}