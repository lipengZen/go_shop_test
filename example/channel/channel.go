package main

import (
	"fmt"
)

func main() {
	queue := make(chan int, 10)
	go func() {
		// rand.Seed(time.Now().Unix())
		queue <- 100 //rand.Intn(10000)

		for {
			// queue <- rand.Intn(10000)
			// time.Sleep(time.Second)
		}

	}()

	for v := range queue {
		// time.Sleep(time.Second * 3)
		fmt.Println(v)
	}
}
