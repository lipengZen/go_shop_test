package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	Message string
}

var (
	queue = make([]Item, 0)
	// cond  = sync.Cond{} 尴尬，初始化弄错了
	cond = sync.NewCond(&sync.Mutex{})
)

func Produce(i int) {

	fmt.Println("i: ", i)
	time.Sleep(time.Second)
	fmt.Println("sleep 1s")

	cond.L.Lock()
	fmt.Println("produce lock")
	defer func() {
		cond.L.Unlock()
		fmt.Println("produce unlock")
	}() // 把这段代码提到 sleep 前面去，就没有问题

	item := Item{
		Message: fmt.Sprintf("new item from : %d", i),
	}
	queue = append(queue, item)
	fmt.Println("add queue: ", queue)

	fmt.Println("send signal")
	cond.Signal()
}

func Consume() {
	cond.L.Lock()
	fmt.Println("consume lock")
	defer func() {
		cond.L.Unlock()
		fmt.Println("consume unlock")
	}()

	if len(queue) <= 0 {
		fmt.Println("wait begin")
		cond.Wait()
		fmt.Println("wait end, queue: ", queue)
	}
	// cond.L.Lock()
	// defer cond.L.Unlock()   原来我真的锁加错了地方
	fmt.Println("queue: ", queue)
	item := queue[0]
	queue = queue[1:]
	fmt.Println(item.Message)

}
func main() {

	var i int
	for {
		i++
		go Produce(i)
		go Consume()

		time.Sleep(time.Second)
	}

}
