package main

import (
	"fmt"
)

const modNum = 1e9 + 7

func fib(n int) int {
	if n == 0 {
		return 0
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		curr, prev = (prev+curr)%modNum, curr
	}

	return curr
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

func test05() {
	n := 95
	succResult := 407059028
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fib(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
