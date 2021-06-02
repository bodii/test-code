package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)
	copyNums := make([]int, numsLen)
	for k, v := range nums {
		copyNums[k] = v
	}

	sort.Ints(copyNums)

	for k, v := range nums {
		result[k] = sort.SearchInts(copyNums, v)
	}

	return result
}

func test01() {
	nums := []int{8, 1, 2, 2, 3}
	succResult := []int{4, 0, 1, 1, 3}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallerNumbersThanCurrent(nums))
	fmt.Println()
}

func test02() {
	nums := []int{6, 5, 4, 8}
	succResult := []int{2, 1, 0, 3}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallerNumbersThanCurrent(nums))
	fmt.Println()
}

func test03() {
	nums := []int{7, 7, 7, 7}
	succResult := []int{0, 0, 0, 0}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallerNumbersThanCurrent(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
