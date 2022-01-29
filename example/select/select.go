package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done channel is triggerred")
				return
			}
		}
	}()

	close(done)

	chan1 := make(chan int)
	chan2 := make(chan int)

	// go func() {
	// 	chan1 <- 1
	// }()
	// go func() {
	// 	chan2 <- 2
	// }()
	timer := time.NewTimer(time.Second)
	select {
	case v := <-chan1:
		println(v)
	case v := <-chan2:
		println(v)
	case <-timer.C:
		fmt.Println("time out")

	}

	chan3 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		chan3 <- 3
	}()
	fmt.Println(<-chan3)

}
