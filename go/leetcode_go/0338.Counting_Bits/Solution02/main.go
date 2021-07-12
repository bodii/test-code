package main

import (
	"fmt"
)

func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		result[i] = result[i/2] + i&1
	}

	return result
}

func test01() {
	n := 2
	succResult := []int{0, 1, 1}
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBits(n))
	fmt.Println()
}

func test02() {
	n := 5
	succResult := []int{0, 1, 1, 2, 1, 2}
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBits(n))
	fmt.Println()
}

func test03() {
	n := 0
	succResult := []int{0}
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBits(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
