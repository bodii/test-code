package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	numsLen := len(nums)
	return max(nums[0]*nums[1]*nums[2], nums[numsLen-1]*nums[numsLen-2]*nums[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	nums := []int{1, 2, 3}
	succResult := 6
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumProduct(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 3, 4}
	succResult := 24
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumProduct(nums))
	fmt.Println()
}

func test03() {
	nums := []int{-1, -2, -3}
	succResult := -6
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumProduct(nums))
	fmt.Println()
}

func test04() {
	nums := []int{-100, -98, -1, 2, 3, 4}
	succResult := 39200
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumProduct(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
