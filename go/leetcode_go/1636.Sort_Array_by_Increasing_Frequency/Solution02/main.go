package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	numsMap := make(map[int]int)

	for _, v := range nums {
		numsMap[v]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if numsMap[nums[i]] == numsMap[nums[j]] {
			return nums[i] > nums[j]
		}
		return numsMap[nums[i]] < numsMap[nums[j]]
	})

	return nums
}

func test01() {
	nums := []int{1, 1, 2, 2, 2, 3}
	succResult := []int{3, 1, 1, 2, 2, 2}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", frequencySort(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 3, 1, 3, 2}
	succResult := []int{1, 3, 3, 2, 2}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", frequencySort(nums))
	fmt.Println()
}

func test03() {
	nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	succResult := []int{5, -1, 4, 4, -6, -6, 1, 1, 1}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", frequencySort(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
