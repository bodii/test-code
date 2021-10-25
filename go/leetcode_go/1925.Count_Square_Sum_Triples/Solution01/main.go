package main

import (
	"fmt"
)

func countTriples(n int) int {
	num := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				if i*i+j*j == k*k {
					num++
				}
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
