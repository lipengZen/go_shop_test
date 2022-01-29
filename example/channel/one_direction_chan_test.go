package main

import (
	"fmt"
	"testing"
	"time"
)

func readChan(myChan <-chan struct{}) {

	fmt.Println("start read: ", len(myChan))
	fmt.Println(<-myChan)

}

func writeChan(myChan chan<- struct{}) {

	fmt.Println("start write: ", len(myChan))
	myChan <- struct{}{}

}

func TestOneDirectionChan(t *testing.T) {

	var myChan chan struct{}
	myChan = make(chan struct{})

	go readChan(myChan)
	time.Sleep(time.Second)
	writeChan(myChan)
}
