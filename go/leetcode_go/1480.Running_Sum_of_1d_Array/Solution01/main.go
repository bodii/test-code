package main

import (
	"fmt"
)

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func test01() {
	nums := []int{1, 2, 3, 4}
	succResult := []int{1, 3, 6, 10}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", runningSum(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 1, 1, 1}
	succResult := []int{1, 2, 3, 4, 5}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", runningSum(nums))
	fmt.Println()
}

func test03() {
	nums := []int{3, 1, 2, 10, 1}
	succResult := []int{3, 4, 6, 16, 17}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", runningSum(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
