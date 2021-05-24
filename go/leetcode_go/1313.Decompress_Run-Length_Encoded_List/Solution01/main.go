package main

import (
	"fmt"
)

func decompressRLElist(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
	}

	return result
}

func test01() {
	nums := []int{1, 2, 3, 4}
	succResult := []int{2, 4, 4, 4}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", decompressRLElist(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 2, 3}
	succResult := []int{1, 3, 3}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", decompressRLElist(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
