package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]bool)

	l, r := 0, len(nums)-1
	for l < r {
		if numsMap[target-nums[l]] {
			return []int{nums[l], target - nums[l]}
		} else {
			numsMap[nums[l]] = true
		}

		if numsMap[target-nums[r]] {
			return []int{nums[r], target - nums[r]}
		} else {
			numsMap[nums[r]] = true
		}
		l++
		r--
	}

	return []int{}
}

func test01() {
	nums := []int{2, 7, 11, 15}
	target := 9
	succResult := []int{2, 7}
	fmt.Println("test01 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoSum(nums, target))
	fmt.Println()
}

func test02() {
	nums := []int{10, 26, 30, 31, 47, 60}
	target := 40
	succResult := []int{10, 30}
	fmt.Println("test02 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoSum(nums, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
