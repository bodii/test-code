package main

import (
	"fmt"
)

func arrangeCoins(n int) int {
	for i := 1; n >= 0; i++ {
		if n < i {
			return i - 1
		}
		n -= i
	}

	return 0
}

func test01() {
	n := 5
	successResult := 2
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", arrangeCoins(n))
	fmt.Println()
}

func test02() {
	n := 8
	successResult := 3
	fmt.Println("test02 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", arrangeCoins(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
