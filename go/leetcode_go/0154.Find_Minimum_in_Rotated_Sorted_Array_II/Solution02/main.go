package main

import (
	"fmt"
	"sort"
)

func findMin(nums []int) int {
	sort.Ints(nums)

	return nums[0]
}

func test01() {
	nums := []int{1, 3, 5}
	succResult := 1
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMin(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 2, 2, 0, 1}
	succResult := 0
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMin(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
