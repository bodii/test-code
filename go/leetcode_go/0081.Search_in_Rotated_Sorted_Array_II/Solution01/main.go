package main

import (
	"fmt"
)

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1

		if nums[mid] == target {
			return true
		}

		if nums[mid] > nums[l] {
			if nums[mid] > target && nums[l] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[l] {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			l++
		}
	}

	return false
}

func test01() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	succResult := true
	fmt.Println("test01 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", search(nums, target))
	fmt.Println()
}

func test02() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 3
	succResult := false
	fmt.Println("test02 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", search(nums, target))
	fmt.Println()
}

func test03() {
	nums := []int{1, 0, 1, 1, 1}
	target := 0
	succResult := true
	fmt.Println("test03 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", search(nums, target))
	fmt.Println()
}

func test04() {
	nums := []int{1, 1, 1, 0, 1}
	target := 0
	succResult := true
	fmt.Println("test04 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", search(nums, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
