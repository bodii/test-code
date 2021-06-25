package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	numsMap := make(map[int]bool)
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if numsMap[nums[i]] {
			result[0] = nums[i]
		}
		numsMap[nums[i]] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !numsMap[i] {
			result[1] = i
			break
		}
	}

	return result
}

func test01() {
	nums := []int{1, 2, 2, 4}
	succResult := []int{2, 3}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1}
	succResult := []int{1, 2}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 2, 3, 5, 4}
	succResult := []int{2, 6}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
