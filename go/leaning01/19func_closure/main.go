package main

import (
	"fmt"
)


// 闭包 colsure
func f1 (f func()) {
	fmt.Println ("The is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("The is f2")
	fmt.Println(x + y)
}


func main() {
	f3 := func () {
		tmp := func(n int, m int) {
			f2(n, m)
		}

		tmp(1, 3)
	}

	f1(f3)

	ret := adder()
	ret2 := ret(200)
	fmt.Println(ret2) // 300
	ret3 := ret(200)
	fmt.Println(ret3) // 500
}


func adder() func(int) int {
	var x int = 100
	return func(y int) int {
		x += y
		return x
	}
}

