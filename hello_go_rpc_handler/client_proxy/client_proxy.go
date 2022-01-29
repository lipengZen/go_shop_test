package client_proxy0

import (
	"net/rpc"
	"oldpackage_go_shop/hello_go_rpc_handler/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceStub(protocol, addr string) HelloServiceStub {

	conn, _ := rpc.Dial(protocol, addr)
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {

	err := c.Client.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
