package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("aaa")
}

type Resource struct {
	url        string
	polliong   bool
	lastPolled int64
}

type Resources struct {
	data []*Resource
	lock *sync.Mutex
}

func Poller(res *Resources) {
	for {
		res.lock.Lock()
		var r *Resource
		for _, v := range res.data {
			if v.polliong {
				continue
			}
			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}
		}
		if r != nil {
			r.polliong = true
		}
		res.lock.Unlock()
		r.polliong = false
		r.lastPolled = time.Now().Unix()
		res.lock.Unlock()
	}
}

//////  利用chan
//type Resource string
//
//func Poller(in,out chan *Resource)  {
//	for i:=range in{
//		out <-i
//	}
//}
