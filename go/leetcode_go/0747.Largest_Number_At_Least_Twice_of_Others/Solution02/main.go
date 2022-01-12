package main

import (
	"fmt"
	"sort"
)

func dominantIndex(nums []int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return 0
	}

	nums2 := make([]int, 101)
	for k, v := range nums {
		nums2[v] = k
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	if nums[0] == 0 || nums[0]/2 < nums[1] {
		return -1
	}

	return nums2[nums[0]]
}

func test01() {
	nums := []int{3, 6, 1, 0}
	succResult := 1
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", dominantIndex(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 3, 4}
	succResult := -1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", dominantIndex(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1}
	succResult := 0
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", dominantIndex(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
