package main

import (
	"fmt"
)

func targetIndices(nums []int, target int) []int {
	startIndex, few := 0, 0
	for _, n := range nums {
		if n < target {
			startIndex++
		} else if n == target {
			few++
		}
	}

	result := make([]int, few)
	for k := range result {
		result[k] = startIndex + k
	}

	return result
}

func test01() {
	nums := []int{1, 2, 5, 2, 3}
	target := 2
	succResult := []int{1, 2}
	fmt.Println("test01 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", targetIndices(nums, target))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 5, 2, 3}
	target := 3
	succResult := []int{3}
	fmt.Println("test02 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", targetIndices(nums, target))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 5, 2, 3}
	target := 5
	succResult := []int{4}
	fmt.Println("test03 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", targetIndices(nums, target))
	fmt.Println()
}

func test04() {
	nums := []int{1, 2, 5, 2, 3}
	target := 4
	succResult := []int{}
	fmt.Println("test04 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", targetIndices(nums, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
