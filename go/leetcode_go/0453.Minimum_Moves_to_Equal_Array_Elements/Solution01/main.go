package main

import (
	"fmt"
	"sort"
)

func minMoves(nums []int) int {
	count := 0

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		count += nums[i] - nums[0]
	}

	return count
}

func test01() {
	nums := []int{1, 2, 3}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMoves(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 1000000000}
	succResult := 999999999
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMoves(nums))
	fmt.Println()
}

func test03() {
	nums := []int{-1000000000, -1}
	succResult := 999999999
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMoves(nums))
	fmt.Println()
}

func test04() {
	nums := []int{5, 6, 8, 8, 5}
	succResult := 7
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMoves(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
