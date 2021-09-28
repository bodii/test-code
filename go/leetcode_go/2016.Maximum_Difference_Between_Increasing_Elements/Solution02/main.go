package main

import (
	"fmt"
)

func maximumDifference(nums []int) int {
	min, maxDiff := nums[0], -1
	for i := 1; i < len(nums); i++ {
		maxDiff = max(maxDiff, nums[i]-min)

		if nums[i] < min {
			min = nums[i]
		}
	}

	return maxDiff
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	nums := []int{7, 1, 5, 4}
	succResult := 4
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumDifference(nums))
	fmt.Println()
}

func test02() {
	nums := []int{9, 4, 3, 2}
	succResult := -1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumDifference(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 5, 2, 10}
	succResult := 9
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumDifference(nums))
	fmt.Println()
}

func test04() {
	nums := []int{0, 0}
	succResult := 0
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumDifference(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
