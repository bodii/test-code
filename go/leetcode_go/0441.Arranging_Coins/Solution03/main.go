package main

import (
	"fmt"
)

func arrangeCoins(n int) int {
	left, right := 0, n
	for left <= right {
		mid := (left + right) / 2
		if (mid+1)*mid/2 <= n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
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
