package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"

	Helloworld "oldpackage_go_shop/hello_go_rpc/proto"
)

func main() {

	req := Helloworld.HelloRequest{
		Name:    "lee",
		Age:     19,
		Classes: []string{"chinese math english art music"},
	}

	rsp, _ := proto.Marshal(&req) // protobuf 具体编码的原理： varint
	fmt.Println(len(rsp))

	bytes, _ := json.Marshal(&req)
	fmt.Println(len(bytes))

}
