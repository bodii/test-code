package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	result := make([]int, n+1)
	result[1], result[2] = 1, 1

	for i := 3; i <= n; i++ {
		result[i] += result[i-1] + result[i-2]
	}

	return result[n]
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
