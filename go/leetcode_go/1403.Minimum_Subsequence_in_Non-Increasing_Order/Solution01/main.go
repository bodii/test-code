package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	numsLen := len(nums)

	leftSum, rightSum := nums[0], 0
	for i := 1; i < numsLen; i++ {
		rightSum += nums[i]
	}

	result := []int{nums[0]}
	for i := 1; i < numsLen; i++ {
		if leftSum > rightSum {
			return result
		}

		leftSum += nums[i]
		rightSum -= nums[i]
		result = append(result, nums[i])
	}

	return result
}

func test01() {
	nums := []int{4, 3, 10, 9, 8}
	succResult := []int{10, 9}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minSubsequence(nums))
	fmt.Println()
}

func test02() {
	nums := []int{4, 4, 7, 6, 7}
	succResult := []int{7, 7, 6}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minSubsequence(nums))
	fmt.Println()
}

func test03() {
	nums := []int{6}
	succResult := []int{6}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minSubsequence(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
