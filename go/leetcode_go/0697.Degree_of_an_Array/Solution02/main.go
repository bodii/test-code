package main

import (
	"fmt"
)

type element struct {
	n, l, r int
}

func findShortestSubArray(nums []int) int {
	maxLen, min := 0, 0
	numMap := make(map[int]*element)
	for k, v := range nums {
		_, ok := numMap[v]
		if !ok {
			numMap[v] = &element{1, k, 0}
		} else {
			numMap[v].n++
			numMap[v].r = k
		}

	}

	for _, v := range numMap {
		if maxLen > v.n {
			continue
		}

		l := v.r - v.l + 1
		if maxLen < v.n {
			maxLen, min = v.n, l
		} else if v.n == maxLen && min > l {
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
