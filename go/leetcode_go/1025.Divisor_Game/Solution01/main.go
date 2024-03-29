package main

import (
	"fmt"
)

func divisorGame(n int) bool {
	return n%2 == 0
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
