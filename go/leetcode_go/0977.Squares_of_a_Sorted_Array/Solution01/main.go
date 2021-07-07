package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	for k, v := range nums {
		nums[k] = v * v
	}
	sort.Ints(nums)

	return nums
}

func test01() {
	nums := []int{-4, -1, 0, 3, 10}
	succResult := []int{0, 1, 9, 16, 100}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortedSquares(nums))
	fmt.Println()
}

func test02() {
	nums := []int{-7, -3, 2, 3, 11}
	succResult := []int{4, 9, 9, 49, 121}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortedSquares(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
