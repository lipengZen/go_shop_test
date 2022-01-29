package main

import (
	"fmt"
	"testing"
	"time"
)

func TestForRange(t *testing.T) {

	myChan := make(chan int, 2)

	go func() {
		for {
			time.Sleep(time.Second)
			myChan <- 1
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		close(myChan)
		fmt.Println("chan closed")
	}()

	for val := range myChan {
		fmt.Println(val)
	}
}
