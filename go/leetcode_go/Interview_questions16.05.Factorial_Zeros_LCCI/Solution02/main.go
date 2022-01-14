package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	count := 0

	for n >= 5 {
		n /= 5
		count += n
	}

	return count
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

func test04() {
	n := 10
	succResult := 2
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trailingZeroes(n))
	fmt.Println()
}

func test05() {
	n := 20
	succResult := 4
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trailingZeroes(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
