package main

import (
	"fmt"
	"net"
	"net/rpc"
	"oldpackage_go_shop/hello_go_rpc_handler/handler"
	"oldpackage_go_shop/hello_go_rpc_handler/server_proxy"
)

func main() {
	// 1. 实例化一个server
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// 2. 注册处理逻辑 handler
	// _ = rpc.RegisterName(handler.HelloServiceName, &handler.HelloService{})
	err = server_proxy.RegisterHelloService(&handler.HelloService{})
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// http 方式
	/*
		http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
			var conn io.ReadWriteCloser = struct {
				io.Writer
				io.ReadCloser
			}{
				ReadCloser: r.Body,
				Writer:     w,
			}
			rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		})

		http.ListenAndServe(":1234", nil)

	*/

	// 3. 启动服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}

	// 改成 新的编解码方式，
	// 很简单：使用了 json
	// rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

}
