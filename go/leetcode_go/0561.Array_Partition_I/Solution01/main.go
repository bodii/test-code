package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	numsLen := len(nums)
	sort.Ints(nums)

	result := 0
	for i := 0; i < numsLen; i += 2 {
		result += nums[i]
	}

	return result
}

func test01() {
	nums := []int{1, 4, 3, 2}
	successResult := 4
	fmt.Println("test01 nums:", nums)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", arrayPairSum(nums))
	fmt.Println()
}

func test02() {
	nums := []int{6, 2, 6, 5, 1, 2}
	successResult := 9
	fmt.Println("test02 nums:", nums)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", arrayPairSum(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
