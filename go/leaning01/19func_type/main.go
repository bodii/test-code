package main

import (
	"fmt"
)

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 10
}

func f3(x func(), y func()int) {
	x()
	fmt.Println(y())
}

func f4(x int, y int)int {
	return x + y
}

func f5(x func()int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}


func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(a, b)

	fmt.Printf("%T\n", f4)

	f6 := f5(f2)
	fmt.Printf("%T\n", f6)
	fmt.Println(f6(1, 3))
}
