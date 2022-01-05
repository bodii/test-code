package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	for i := 0; i < k; i++ {
		nums[0] = -nums[0]

		if nums[0] > nums[1] {
			for i := 1; i < len(nums); i++ {
				if nums[i-1] > nums[i] {
					nums[i-1], nums[i] = nums[i], nums[i-1]
				} else {
					break
				}
			}
		}

	}

	result := 0
	for _, v := range nums {
		result += v
	}

	return result
}

func test01() {
	nums := []int{4, 2, 3}
	k := 1
	succResult := 5
	fmt.Println("test01 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestSumAfterKNegations(nums, k))
	fmt.Println()
}

func test02() {
	nums := []int{3, -1, 0, 2}
	k := 3
	succResult := 6
	fmt.Println("test02 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestSumAfterKNegations(nums, k))
	fmt.Println()
}

func test03() {
	nums := []int{2, -3, -1, 5, -4}
	k := 2
	succResult := 13
	fmt.Println("test03 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestSumAfterKNegations(nums, k))
	fmt.Println()
}

func test04() {
	nums := []int{2, -3, -1, 5, -4}
	k := 7
	succResult := 15
	fmt.Println("test04 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestSumAfterKNegations(nums, k))
	fmt.Println()
}

func test05() {
	nums := []int{-2, 5, 0, 2, -2}
	k := 3
	succResult := 11
	fmt.Println("test05 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestSumAfterKNegations(nums, k))
	fmt.Println()
}

func test06() {
	nums := []int{-6, -8, -7, 1}
	k := 4
	succResult := 20
	fmt.Println("test06 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestSumAfterKNegations(nums, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
