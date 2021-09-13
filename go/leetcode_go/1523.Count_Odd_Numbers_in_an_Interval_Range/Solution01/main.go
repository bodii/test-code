package main

import (
	"fmt"
)

func countOdds(low int, high int) int {
	if low%2 == 0 && high%2 == 0 {
		return (high - low) / 2
	} else {
		return (high-low)/2 + 1
	}
}

func test01() {
	low, high := 3, 7
	succResult := 3
	fmt.Println("test01 low:=", low, " high:=", high)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countOdds(low, high))
	fmt.Println()
}

func test02() {
	low, high := 8, 10
	succResult := 1
	fmt.Println("test02 low:=", low, " high:=", high)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countOdds(low, high))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
