package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	numsLen := len(nums)
	result := []int{}
	for i := 0; i < numsLen; i++ {
		index := (nums[i] - 1) % numsLen
		nums[index] += numsLen
	}

	for i := 0; i < numsLen; i++ {
		if nums[i] <= numsLen {
			result = append(result, i+1)
		}
	}

	return result
}

func test01() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	succResult := []int{5, 6}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findDisappearedNumbers(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1}
	succResult := []int{2}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findDisappearedNumbers(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 7, 7, 8, 9, 9, 9, 9, 10}
	succResult := []int{4, 5, 6, 11, 12, 13, 14, 15}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findDisappearedNumbers(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
