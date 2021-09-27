package main

import (
	"fmt"
)

func findMiddleIndex(nums []int) int {
	l, r := 0, 0
	for _, v := range nums {
		r += v
	}

	for i := 0; i < len(nums); i++ {
		r -= nums[i]
		if l == r {
			return i
		}
		l += nums[i]
	}

	return -1
}

func test01() {
	nums := []int{2, 3, -1, 8, 4}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMiddleIndex(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, -1, 4}
	succResult := 2
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMiddleIndex(nums))
	fmt.Println()
}

func test03() {
	nums := []int{2, 5}
	succResult := -1
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMiddleIndex(nums))
	fmt.Println()
}

func test04() {
	nums := []int{1}
	succResult := 0
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMiddleIndex(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
