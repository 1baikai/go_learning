package main

import "fmt"

func f(x int) (y int) {
	fmt.Println("x", x)
	defer func() {
		x++
		fmt.Println("x", x)
	}()
	fmt.Println("X", x)
	return x
}

func f1(x int) int {

	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // y =x =5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(f1(1))
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
