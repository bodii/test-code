package main

import (
	"fmt"
	"sort"
)

func exchange(nums []int) []int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i]%2 == 1 && nums[j]%2 == 0
	})

	return nums
}

func test01() {
	nums := []int{1, 2, 3, 4}
	succResult := []int{1, 3, 2, 4}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", exchange(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 2, 3}
	succResult := []int{1, 1, 3, 2}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", exchange(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
