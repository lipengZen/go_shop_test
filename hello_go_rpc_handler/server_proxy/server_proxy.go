package server_proxy

import (
	"net/rpc"
	"oldpackage_go_shop/hello_go_rpc_handler/handler"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

// 这里耦合严重， 和 handler.HelloService 强耦合
// 关心的是结构体的函数，不是结构体名称

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
