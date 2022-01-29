package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"oldpackage_go_shop/hello_grpc_stream/proto"
)

const PORT = ":50052"

type Server struct {
}

func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		str := strconv.Itoa(i)
		res.Send(&proto.StreamResData{Data: str})
		i++
		time.Sleep(1 * time.Second)
		if i > 20 {
			break
		}
	}
	return nil
}

func (s *Server) PutStream(cliStr proto.Greeter_PutStreamServer) error {

	for {
		a, err := cliStr.Recv()
		if err != nil {
			fmt.Println("error : ", err)
			return err
		}
		fmt.Println("a : ", a)
	}
	// return nil
}
func (s *Server) AllStream(allStr proto.Greeter_AllStreamServer) error {

	for {
		a, err := allStr.Recv()
		if err != nil {
			fmt.Println("error : ", err)
			return err
		}
		fmt.Println("xiaopengzi : ", a.Data)

		allStr.Send(&proto.StreamResData{Data: "嘻嘻"})
	}

	// return nil
}

func main() {

	g := grpc.NewServer()

	proto.RegisterGreeterServer(g, &Server{})

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("failed to listen: ", err)
	}

	g.Serve(lis)

}
