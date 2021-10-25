package main

import (
	"fmt"
)

func countTriples(n int) int {
	num := 0
	sum := make(map[int]bool)
	square := make([]int, n+1)
	for i := 1; i <= n; i++ {
		square[i] = i * i
		sum[square[i]] = true
	}

	for i := 3; i <= n; i++ {
		for j := 3; j <= n; j++ {
			if sum[square[i]+square[j]] {
				// fmt.Println(i, j, sum[sqrt[i]+sqrt[j]])
				num++
			}
		}
	}

	return num
}

func test01() {
	n := 5
	successResult := 2
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", countTriples(n))
	fmt.Println()
}

func test02() {
	n := 10
	successResult := 4
	fmt.Println("test02 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", countTriples(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
