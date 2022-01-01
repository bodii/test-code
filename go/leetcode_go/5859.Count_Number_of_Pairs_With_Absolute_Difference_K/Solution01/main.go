package main

import (
	"fmt"
)

func countKDifference(nums []int, k int) int {
	numsLen := len(nums)
	count := 0

	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if abs(nums[i]-nums[j]) == k {
				count++
			}
		}
	}

	return count
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func test01() {
	nums := []int{1, 2, 2, 1}
	k := 1
	succResult := 4
	fmt.Println("test01 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countKDifference(nums, k))
	fmt.Println()
}

func test02() {
	nums := []int{1, 3}
	k := 3
	succResult := 0
	fmt.Println("test02 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countKDifference(nums, k))
	fmt.Println()
}

func test03() {
	nums := []int{3, 2, 1, 5, 4}
	k := 2
	succResult := 3
	fmt.Println("test03 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countKDifference(nums, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
