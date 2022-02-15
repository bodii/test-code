package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 0 {
		return 0
	}

	list := make([]int, n+1)
	list[0], list[1] = 0, 1
	for i := 2; i <= n; i++ {
		list[i] = list[i-1] + list[i-2]
	}

	return list[n] % 1000000007
}

func test01() {
	n := 2
	succResult := 1
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fib(n))
	fmt.Println()
}

func test02() {
	n := 5
	succResult := 5
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fib(n))
	fmt.Println()
}

func test03() {
	n := 30
	succResult := 832040
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fib(n))
	fmt.Println()
}

func test04() {
	n := 45
	succResult := 134903163
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fib(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
