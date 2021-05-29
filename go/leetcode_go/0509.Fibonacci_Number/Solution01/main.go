package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func test01() {
	n := 2
	successResult := 1
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", fib(n))
	fmt.Println()
}

func test02() {
	n := 3
	successResult := 2
	fmt.Println("test02 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", fib(n))
	fmt.Println()
}

func test03() {
	n := 4
	successResult := 3
	fmt.Println("test03 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", fib(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
