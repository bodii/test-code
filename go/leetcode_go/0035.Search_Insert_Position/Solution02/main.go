package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	len := len(nums)
	if nums[0] > target {
		return 0
	}

	if nums[0] <= target && nums[len-1] >= target {
		return dichotomy(nums, 0, len-1, target)
	}

	return len
}

func dichotomy(nums []int, start, end, target int) int {
	if start > end {
		return start
	}

	mid := start + (end-start)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return dichotomy(nums, start, mid-1, target)
	}

	return dichotomy(nums, mid+1, end, target)
}

func test01() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println("test01 is 2 result:", searchInsert(nums, target))
}

func test02() {
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println("test02 is 1 result:", searchInsert(nums, target))
}

func test03() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println("test03 is 4 result:", searchInsert(nums, target))
}

func test04() {
	nums := []int{1, 3, 5, 6}
	target := 0
	fmt.Println("test04 is 0 result:", searchInsert(nums, target))
}

func test05() {
	nums := []int{1, 3}
	target := 2
	fmt.Println("test05 is 1 result:", searchInsert(nums, target))
}

func test06() {
	nums := []int{1, 3}
	target := 1
	fmt.Println("test06 is 0 result:", searchInsert(nums, target))
}

func test07() {
	nums := []int{1, 3}
	target := 4
	fmt.Println("test07 is 2 result:", searchInsert(nums, target))
}

func test08() {
	nums := []int{1, 3}
	target := 3
	fmt.Println("test08 is 1 result:", searchInsert(nums, target))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
}
