package main

import (
	"fmt"
)

func sumOfDigits(nums []int) int {
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		minNum = min(minNum, nums[i])
	}

	sum := 0
	for ; minNum > 0; minNum /= 10 {
		sum += minNum % 10
	}

	if sum&1 == 0 {
		return 1
	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func test01() {
	nums := []int{34, 23, 1, 24, 75, 33, 54, 8}
	succResult := 0
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumOfDigits(nums))
	fmt.Println()
}

func test02() {
	nums := []int{99, 77, 33, 66, 55}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumOfDigits(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
