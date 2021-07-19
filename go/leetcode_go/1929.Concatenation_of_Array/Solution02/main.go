package main

import (
	"fmt"
)

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func test01() {
	nums := []int{1, 2, 1}
	succResult := []int{1, 2, 1, 1, 2, 1}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getConcatenation(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 3, 2, 1}
	succResult := []int{1, 3, 2, 1, 1, 3, 2, 1}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getConcatenation(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
