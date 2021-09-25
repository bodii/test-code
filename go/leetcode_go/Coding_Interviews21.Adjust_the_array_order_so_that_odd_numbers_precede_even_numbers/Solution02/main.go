package main

import (
	"fmt"
)

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]%2 == 0 && nums[right]%2 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else if nums[left]%2 == 1 {
			left++
		} else if nums[right]%2 == 0 {
			right--
		}
	}

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
