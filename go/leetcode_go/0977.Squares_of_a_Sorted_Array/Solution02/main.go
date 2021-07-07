package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	numsLen := len(nums)
	left, right, index := 0, numsLen-1, numsLen-1

	result := make([]int, numsLen)
	for left <= right {
		leftN := nums[left] * nums[left]
		rightN := nums[right] * nums[right]

		if leftN > rightN {
			result[index] = leftN
			left++
		} else {
			result[index] = rightN
			right--
		}
		index--
	}

	return result
}

func test01() {
	nums := []int{-4, -1, 0, 3, 10}
	succResult := []int{0, 1, 9, 16, 100}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortedSquares(nums))
	fmt.Println()
}

func test02() {
	nums := []int{-7, -3, 2, 3, 11}
	succResult := []int{4, 9, 9, 49, 121}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortedSquares(nums))
	fmt.Println()
}

func test03() {
	nums := []int{-4, 0, 2, 5, 7}
	succResult := []int{0, 4, 16, 25, 49}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortedSquares(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
