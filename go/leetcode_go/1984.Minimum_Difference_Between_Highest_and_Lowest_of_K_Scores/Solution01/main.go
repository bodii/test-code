package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	numsLen := len(nums)

	if numsLen == 1 {
		return 0
	}

	minDiff := math.MaxInt32
	sort.Ints(nums)
	for i := k - 1; i < numsLen; i++ {
		minDiff = min(minDiff, nums[i]-nums[i-k+1])
	}

	return minDiff
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func test01() {
	nums := []int{90}
	k := 1
	succResult := 0
	fmt.Println("test01 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumDifference(nums, k))
	fmt.Println()
}

func test02() {
	nums := []int{9, 4, 1, 7}
	k := 2
	succResult := 2
	fmt.Println("test02 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumDifference(nums, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
