package main

import (
	"fmt"
)

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}

	return nums[l]
}

func test01() {
	nums := []int{1, 3, 5}
	succResult := 1
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMin(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 2, 2, 0, 1}
	succResult := 0
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMin(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
