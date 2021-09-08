package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func test01() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	succResult := 4
	fmt.Println("test01 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", search(nums, target))
	fmt.Println()
}

func test02() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	succResult := -1
	fmt.Println("test02 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", search(nums, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
