package main

import (
	"fmt"
)

func findLengthOfLCIS(nums []int) int {
	numsLen := len(nums)

	count := 1
	for i := 0; i < numsLen-1; i++ {
		if nums[i] >= nums[i+1] {
			continue
		}

		j := i + 1
		for ; j < numsLen; j++ {
			if nums[j] <= nums[j-1] {
				break
			}
		}

		count = max(count, j-i)
		i = j - 1
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	nums := []int{1, 3, 5, 4, 7}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findLengthOfLCIS(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 2, 2, 2, 2}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findLengthOfLCIS(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
