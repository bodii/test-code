package main

import (
	"fmt"
)

func sortArrayByParityII(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)

	l1, l2 := 0, 1
	for i := 0; i < numsLen; i++ {
		if nums[i]%2 == 0 {
			result[l1] = nums[i]
			l1 += 2
		} else {
			result[l2] = nums[i]
			l2 += 2
		}
	}

	return result
}

func test01() {
	nums := []int{4, 2, 5, 7}
	success := []int{4, 5, 2, 7}

	fmt.Println("test01  nums:=", nums)
	fmt.Println("success result:", success)
	fmt.Println("result:", sortArrayByParityII(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 3}
	success := []int{2, 3}

	fmt.Println("test02  nums:=", nums)
	fmt.Println("success result:", success)
	fmt.Println("result:", sortArrayByParityII(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
