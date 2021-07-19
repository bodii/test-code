package main

import (
	"fmt"
)

func buildArray(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = nums[nums[i]]
	}

	return result
}

func test01() {
	nums := []int{0, 2, 1, 5, 3, 4}
	succResult := []int{0, 1, 2, 4, 5, 3}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buildArray(nums))
	fmt.Println()
}

func test02() {
	nums := []int{5, 0, 1, 2, 3, 4}
	succResult := []int{4, 5, 0, 1, 2, 3}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buildArray(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
