package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	len := len(nums)
	if len <= 2 {
		return nums[0]
	}

	sort.Ints(nums)
	for i := 2; i < len; i++ {
		if nums[i-2] != nums[i-1] && i == 2 {
			return nums[i-2]
		}

		if nums[i-1] != nums[i] && i == len-1 {
			return nums[i]
		}

		if nums[i-2] != nums[i-1] && nums[i-1] != nums[i] {
			return nums[i-1]
		}
	}

	return -1
}

func test01() {
	nums := []int{3, 4, 3, 3}
	succResult := 4
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", singleNumber(nums))
	fmt.Println()
}

func test02() {
	nums := []int{9, 1, 7, 9, 7, 9, 7}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", singleNumber(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
