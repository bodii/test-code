package main

import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	d, result := x^y, 0

	for d > 0 {
		result += d & 1
		d >>= 1
	}

	return result
}

func test01() {
	x, y := 1, 4
	succResult := 2
	fmt.Println("test01 x:=", x, " y:=", y)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hammingDistance(x, y))
	fmt.Println()
}

func test02() {
	x, y := 3, 1
	succResult := 1
	fmt.Println("test02 x:=", x, " y:=", y)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hammingDistance(x, y))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
