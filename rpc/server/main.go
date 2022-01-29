package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	// http://127.0.0.1:8000/add?a=1&b=2
	// 返回格式： json {"data":3}

	http.HandleFunc("/add", func(rw http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		fmt.Println("path:", r.URL.Path)
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		rw.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = rw.Write(jData)
	})

	http.ListenAndServe(":8000", nil)
}
