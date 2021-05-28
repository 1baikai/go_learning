package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	//"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

//func main() {
//	MemoryLeak()
//	//reversal()
//	//Concurrent()
//}

// 强制转换会引起内存拷贝
//func main() {
//	a := "aaa"
//	b := *(*reflect.StringHeader)(unsafe.Pointer(&a))
//	fmt.Println("b", b)
//	c := *(*[]byte)(unsafe.Pointer(&b))
//	fmt.Println("C", c)
//}

// MemoryLeak 内存泄露的题，问：再不使用resp.body.close 会产生几个内存泄露 ---3
func MemoryLeak() {
	num := 6
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.baidu.com")
		_, _ = ioutil.ReadAll(resp.Body)

	}

	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}

// 翻转所有的中英数字
func reversal() {
	a := []rune("手机123aadc")
	fmt.Println(a)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(string(a))
}

// Concurrent 并发知识
func Concurrent() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello world")
		mu.Unlock()
	}()
	mu.Lock()
}



func test() {

	log.Println(" ===> loop begin.")
	for i := 0; i < 10; i++ {
		log.Println(genSomeBytes())
	}

	log.Println(" ===> loop end.")
}

//生成一个随机字符串
func genSomeBytes() *bytes.Buffer {

	var buff bytes.Buffer

	for i := 1; i < 10; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}

	return &buff
}

func main() {

	go func() {
		for {
			test()
			time.Sleep(time.Second * 1)
		}
	}()

	//启动pprof
	http.ListenAndServe("0.0.0.0:9090", nil)

}