package main

import (
	"fmt"
)

func minOperations(nums []int) int {
	res := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			increment := nums[i] - nums[i+1] + 1
			res += increment
			nums[i+1] += increment
		}
	}

	return res
}

func test01() {
	nums := []int{1, 1, 1}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 5, 2, 4, 1}
	succResult := 14
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(nums))
	fmt.Println()
}

func test03() {
	nums := []int{8}
	succResult := 0
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(nums))
	fmt.Println()
}

func test04() {
	nums := []int{7, 4, 2, 8, 1, 7, 7, 10}
	succResult := 38
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
