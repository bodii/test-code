package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	prev, cur := 0, 1
	for i := 2; i <= n; i++ {
		tmp := prev
		prev = cur
		cur = tmp + prev
	}

	return cur
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

func test04() {
	n := 5
	successResult := 5
	fmt.Println("test04 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", fib(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
