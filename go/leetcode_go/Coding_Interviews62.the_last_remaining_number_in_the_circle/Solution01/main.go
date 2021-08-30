package main

import (
	"fmt"
)

func lastRemaining(n int, m int) int {
	ns := make([]int, n)
	for i := 0; i < n; i++ {
		ns[i] = i
	}

	for ; n > 1; n-- {
		p := m % n

		if p >= 1 {
			ns = append(ns[0:p-1], ns[p:]...)
		} else {
			ns = ns[0 : n-1]
		}
	}

	return ns[0]
}

func test01() {
	n, m := 5, 3
	succResult := 3
	fmt.Println("test01 n:=", n, " m:=", m)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lastRemaining(n, m))
	fmt.Println()
}

func test02() {
	n, m := 10, 17
	succResult := 2
	fmt.Println("test02 n:=", n, " m:=", m)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lastRemaining(n, m))
	fmt.Println()
}

func test03() {
	n, m := 88, 10
	succResult := 3
	fmt.Println("test03 n:=", n, " m:=", m)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lastRemaining(n, m))
	fmt.Println()
}

func test04() {
	n, m := 71989, 111059
	succResult := 34203
	fmt.Println("test04 n:=", n, " m:=", m)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lastRemaining(n, m))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
