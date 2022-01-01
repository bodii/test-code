package main

import (
	"fmt"
)

func countKDifference(nums []int, k int) int {
	count := 0
	numsMap := make(map[int]int)

	for _, v := range nums {
		numsMap[v]++
		count += numsMap[v+k] + numsMap[v-k]
	}

	return count
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func test01() {
	nums := []int{1, 2, 2, 1}
	k := 1
	succResult := 4
	fmt.Println("test01 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countKDifference(nums, k))
	fmt.Println()
}

func test02() {
	nums := []int{1, 3}
	k := 3
	succResult := 0
	fmt.Println("test02 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countKDifference(nums, k))
	fmt.Println()
}

func test03() {
	nums := []int{3, 2, 1, 5, 4}
	k := 2
	succResult := 3
	fmt.Println("test03 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countKDifference(nums, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
