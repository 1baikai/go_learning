package main

import (
	"fmt"
	"sync"
)

var l sync.Mutex
var a1 string

func main() {

	hello1()
	l.Lock()
	fmt.Println(a1)
}
func hello1() {
	l.Lock()
	go func() {
		a1 = "hello\n"
		l.Unlock()
	}()
}
