package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}

	return n/5 + trailingZeroes(n/5)
}

func test01() {
	n := 3
	succResult := 0
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trailingZeroes(n))
	fmt.Println()
}

func test02() {
	n := 5
	succResult := 1
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trailingZeroes(n))
	fmt.Println()
}

func test03() {
	n := 7
	succResult := 1
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trailingZeroes(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
