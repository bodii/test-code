package main

import (
	"fmt"
)

func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	if numsLen == 1 {
		return nums[0]
	}

	d1, d2 := nums[0], max(nums[0], nums[1])
	for i := 2; i < numsLen; i++ {
		d2, d1 = d1+nums[i], max(d1, d2)
	}

	return max(d1, d2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	nums := []int{1, 2, 3, 1}
	succResult := 4
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", rob(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 7, 9, 3, 1}
	succResult := 12
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", rob(nums))
	fmt.Println()
}

func test03() {
	nums := []int{2, 1, 4, 5, 3, 1, 1, 3}
	succResult := 12
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", rob(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
