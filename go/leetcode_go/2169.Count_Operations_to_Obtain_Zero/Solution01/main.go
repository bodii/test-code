package main

import (
	"fmt"
)

func countOperations(num1 int, num2 int) int {
	step := 0

	for num1 > 0 && num2 > 0 {
		if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		step++
	}

	return step
}

func test01() {
	num1, num2 := 2, 3
	succResult := 3
	fmt.Println("test01 num1:=", num1, " num2:=", num2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countOperations(num1, num2))
	fmt.Println()
}

func test02() {
	num1, num2 := 10, 10
	succResult := 1
	fmt.Println("test02 num1:=", num1, " num2:=", num2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countOperations(num1, num2))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
