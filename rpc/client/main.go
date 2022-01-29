package main

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {

	// http://127.0.0.1:8000/add?a=1&b=2
	// 返回格式： json {"data":3}

	req := HttpRequest.NewRequest()

	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))

	body, _ := res.Body()
	fmt.Println("res: ", string(body))

	rspData := ResponseData{}

	_ = json.Unmarshal(body, &rspData)

	return rspData.Data

}
func main() {

	fmt.Println(" a+b: ", Add(2, 2))

}
