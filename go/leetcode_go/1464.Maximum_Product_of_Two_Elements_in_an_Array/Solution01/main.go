package main

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	sort.Ints(nums)

	numsLenIndex := len(nums) - 1
	return (nums[numsLenIndex] - 1) * (nums[numsLenIndex-1] - 1)
}

func test01() {
	nums := []int{3, 4, 5, 2}
	succResult := 12
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProduct(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 5, 4, 5}
	succResult := 16
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProduct(nums))
	fmt.Println()
}

func test03() {
	nums := []int{3, 7}
	succResult := 12
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProduct(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
