package main

import (
	"fmt"
)

func swapNumbers(numbers []int) []int {
	return []int{numbers[1], numbers[0]}
}

func test01() {
	nums := []int{1, 2}
	succResult := []int{2, 1}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", swapNumbers(nums))
	fmt.Println()
}

func main() {
	test01()
}
