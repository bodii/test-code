package main

import (
	"fmt"
)

func sumZero(n int) []int {
	result := make([]int, n)

	if n%2 != 0 {
		n--
	}

	for i := 1; i <= n; i += 2 {
		result[i-1] = i
		result[i] = -i
	}

	return result
}

func test01() {
	n := 5
	succResult := "[]int{-7,-1,1,3,4}, []int{-5,-1,1,2,3}, []int{-3,-1,2,-2,4} 其中一个"
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumZero(n))
	fmt.Println()
}

func test02() {
	n := 3
	succResult := "[]int{-1,0,1}"
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumZero(n))
	fmt.Println()
}

func test03() {
	n := 1
	succResult := "[]int{0}"
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumZero(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
