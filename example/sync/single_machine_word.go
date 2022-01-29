package main

import "fmt"

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	Id   int
	Name string
}

func (b *Ben) Hello() {
	fmt.Printf("hello, I am %s", b.Name)
}

type Jerry struct {
	Name string
}

func (j *Jerry) Hello() {
	fmt.Printf("hello, I am %s", j.Name)
}

func main() {

	var ben = &Ben{Name: "ben"}

	var jerry = &Jerry{Name: "jerry"}

	var loop0, loop1 func()

	var maker IceCreamMaker = ben

	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}

}
