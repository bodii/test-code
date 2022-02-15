package main

import (
	"fmt"
)

func fib(n int) int {
	list := make([]int, n+1)
	return helper(list, n)
}

func helper(list []int, i int) int {
	if i <= 1 {
		return i
	}

	if list[i] > 0 {
		return list[i]
	}

	list[i] = helper(list, i-1) + helper(list, i-2)
	return list[i] % 1000000007
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
