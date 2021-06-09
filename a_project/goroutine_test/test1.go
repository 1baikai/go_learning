package main

var a string

func main() {
	c := make(chan int, 10)
	hello(c)
	<-c
	print(a)
}

func hello(c chan int) {

	go func() {
		a = "hello\n"
		c <- 0
	}()

}
