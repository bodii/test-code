package main

import (
	"fmt"
)

func sumNums(n int) int {
	var sum func(a int) bool
	nums := 0
	sum = func(a int) bool {
		nums += a
		return a != 0 && sum(a-1)
	}

	sum(n)

	return nums
}

func test01() {
	n := 3
	successResult := 6
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", sumNums(n))
	fmt.Println()
}

func test02() {
	n := 9
	successResult := 45
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", sumNums(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
