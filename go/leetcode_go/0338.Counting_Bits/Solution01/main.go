package main

import (
	"fmt"
)

func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		result[i] = countOneBit(i)
	}

	return result
}

func countOneBit(n int) int {
	result := 0
	for n > 0 {
		n &= (n - 1)
		result++
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

func main() {
	test01()
	test02()
}
