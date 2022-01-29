package main

import "testing"

var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world"
	<-c
}

func TestHappensBefore(t *testing.T) {

	go f()
	c <- 0
	t.Log(a)

}
