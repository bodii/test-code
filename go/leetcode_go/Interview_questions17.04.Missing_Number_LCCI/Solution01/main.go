package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)

	left, right := 0, len(nums)
	for left < right {
		if nums[left] != left {
			return left
		}
		left++
		if nums[right-1] != right {
			return right
		}
		right--
	}

	return right
}

func test01() {
	nums := []int{3, 0, 1}
	succResult := 2
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", missingNumber(nums))
	fmt.Println()
}

func test02() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	succResult := 8
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", missingNumber(nums))
	fmt.Println()
}

func test03() {
	nums := []int{4, 2, 5, 0, 1}
	succResult := 3
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", missingNumber(nums))
	fmt.Println()
}

func test04() {
	nums := []int{4, 2, 5, 0, 1, 6}
	succResult := 3
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", missingNumber(nums))
	fmt.Println()
}

func test05() {
	nums := []int{8, 6, 3, 2, 5, 7, 0, 1}
	succResult := 4
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", missingNumber(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
