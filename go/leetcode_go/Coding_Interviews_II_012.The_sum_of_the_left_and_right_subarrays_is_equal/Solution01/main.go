package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	l, r := 0, 0
	for _, n := range nums {
		r += n
	}

	for k, n := range nums {
		r -= n
		if l == r {
			return k
		}

		l += n
	}

	return -1
}

func test01() {
	nums := []int{1, 7, 3, 6, 5, 6}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", pivotIndex(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 3}
	succResult := -1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", pivotIndex(nums))
	fmt.Println()
}

func test03() {
	nums := []int{2, 1, -1}
	succResult := 0
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", pivotIndex(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
