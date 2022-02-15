package main

import (
	"fmt"
)

func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return (fib(n-1) + fib(n-2)) % 1000000007
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
