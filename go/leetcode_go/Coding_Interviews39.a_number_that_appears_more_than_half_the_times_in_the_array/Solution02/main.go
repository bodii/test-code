package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func test01() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	succResult := 2
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", majorityElement(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", majorityElement(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
