package main

import (
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var counter int = 0

func TestRace(t *testing.T) {

	wg = sync.WaitGroup{}
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go AddValue()
	}

	wg.Wait()

	t.Log("counter: ", counter)
}

func AddValue() {

	for count := 0; count < 20000; count++ {
		value := counter
		time.Sleep(time.Nanosecond * 1)
		value++
		counter = value
	}
	wg.Done()
}
