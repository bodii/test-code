package main

import (
	"fmt"
)

func findComplement(num int) int {
	num1, num2 := num, 1
	for num1 > 0 {
		num1 >>= 1
		num2 <<= 1
	}
	num2 -= 1
	return num ^ num2
}

func test01() {
	num := 5
	successResult := 2
	fmt.Println("test01 num:", num)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", findComplement(num))
	fmt.Println()
}

func test02() {
	num := 1
	successResult := 0
	fmt.Println("test02 num:", num)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", findComplement(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
