package main

import (
	"fmt"
)

func addDigits(num int) int {
	result := num % 9

	if num != 0 && result == 0 {
		return 9
	}

	return result
}

func test01() {
	num := 38
	succResult := 2
	fmt.Println("test01 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addDigits(num))
	fmt.Println()
}

func test02() {
	num := 77
	succResult := 5
	fmt.Println("test02 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addDigits(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
