package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	//a := 123
	//
	//c1 := make(chan error, 1)
	//go func() {
	//	err := aaa(a)
	//	if err != nil {
	//		c1 <- err
	//	}
	//	c1 <- nil
	//}()
	//
	//select {
	//case res := <-c1:
	//	if res!=nil{
	//		fmt.Println(res)
	//	}
	//case <-time.After(time.Second * 10):
	//	fmt.Println("timeout 1")
	//}


	func1()
}

func aaa(a int) error {
	if a/2 == 0 {
		return errors.New("Err")
	}
	return nil
}
func func1() {
	hctx, hcancel := context.WithTimeout(context.Background(), time.Second*4)
	defer hcancel()

	errChan := make(chan error, 1)
	// 处理逻辑
	go func() {
		// 处理耗时
		time.Sleep(time.Second * 6)
		errChan  <- errors.New("wanshi")
	}()

	// 超时机制
	select {

	case <-hctx.Done():
		fmt.Println(hctx.Err(),"hctx timeout")
	case v := <-errChan:
		if v!=nil{
			fmt.Println("youcuowu ")
			return
		}
	}

	return

}