package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	left, right := 0, 0
	for _, v := range nums {
		right += v
	}

	right -= nums[0]
	if left == right {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		right -= nums[i]
		left += nums[i-1]

		if left == right {
			return i
		}
	}

	return -1
}

func test01() {
	nums := []int{1, 7, 3, 6, 5, 6}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", pivotIndex(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 3}
	succResult := -1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", pivotIndex(nums))
	fmt.Println()
}

func test03() {
	nums := []int{2, 1, -1}
	succResult := 0
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", pivotIndex(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
