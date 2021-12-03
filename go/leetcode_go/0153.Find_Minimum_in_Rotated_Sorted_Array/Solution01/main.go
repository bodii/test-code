package main

import (
	"fmt"
)

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)>>1

		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}

func test01() {
	nums := []int{3, 4, 5, 1, 2}
	succResult := 1
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMin(nums))
	fmt.Println()
}

func test02() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	succResult := 0
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMin(nums))
	fmt.Println()
}

func test03() {
	nums := []int{11, 13, 15, 17}
	succResult := 11
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMin(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
