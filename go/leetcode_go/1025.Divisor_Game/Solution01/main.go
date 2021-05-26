package main

import (
	"fmt"
)

func divisorGame(n int) bool {
	x := n - 1
	nums := 0
	for n > 0 && x > 0 {
		if n%x == 0 {
			nums++
		}
		n -= x
		x--
	}

	fmt.Println(nums)

	return nums%2 == 1
}

func test01() {
	n := 2
	successResult := true
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", divisorGame(n))
	fmt.Println()
}

func test02() {
	n := 3
	successResult := false
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", divisorGame(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
