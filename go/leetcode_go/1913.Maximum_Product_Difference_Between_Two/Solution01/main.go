package main

import (
	"fmt"
	"sort"
)

func maxProductDifference(nums []int) int {
	sort.Ints(nums)

	numsLen := len(nums)
	return nums[numsLen-1]*nums[numsLen-2] - nums[0]*nums[1]
}

func test01() {
	nums := []int{5, 6, 2, 7, 4}
	succResult := 34
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProductDifference(nums))
	fmt.Println()
}

func test02() {
	nums := []int{4, 2, 5, 9, 7, 4, 8}
	succResult := 64
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProductDifference(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
