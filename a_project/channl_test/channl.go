package main

import (
	"fmt"
	"sync"
)


// 典型例子
// time  中的 tick 只能取值

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	wg.Add(2)
	go pull(ch1)
	go get(ch1, ch2)
	wg.Wait()
	for y := range ch2 {
		fmt.Printf("取出来的值%d \n", y)
	}

}

func get(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	//for x := range ch1 {
	//	ch2 <- x * x
	//}
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x

	}
	close(ch2)

}

func pull(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}
