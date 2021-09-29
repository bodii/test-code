package main

import (
	"fmt"
)

func findShortestSubArray(nums []int) int {
	maxLen, min := 0, 0
	numMap := make(map[int][]int)
	for k, v := range nums {
		if _, ok := numMap[v]; !ok {
			numMap[v] = make([]int, 3)
			numMap[v][1] = k
		}
		numMap[v][0] += 1
		numMap[v][2] = k
		if maxLen < numMap[v][0] {
			maxLen = numMap[v][0]
		}
	}

	for _, v := range numMap {
		if v[0] < maxLen {
			continue
		}

		l := v[2] - v[1] + 1
		if min == 0 || min > l {
			min = l
		}

	}

	return min
}

func test01() {
	nums := []int{1, 2, 2, 3, 1}
	succResult := 2
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findShortestSubArray(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	succResult := 6
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findShortestSubArray(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1}
	succResult := 1
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findShortestSubArray(nums))
	fmt.Println()
}

func test04() {
	nums := []int{1, 2, 1}
	succResult := 3
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findShortestSubArray(nums))
	fmt.Println()
}

func test05() {
	nums := []int{1, 1, 2, 2, 2, 1}
	succResult := 3
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findShortestSubArray(nums))
	fmt.Println()
}

func test06() {
	nums := []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2}
	succResult := 7
	fmt.Println("test06 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findShortestSubArray(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
