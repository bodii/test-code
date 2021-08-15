package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l, r}
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}

	return []int{}
}

func test01() {
	nums := []int{1, 2, 4, 6, 10}
	target := 8
	succResult := []int{1, 3}
	fmt.Println("test01 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoSum(nums, target))
	fmt.Println()
}

func test02() {
	nums := []int{2, 3, 4}
	target := 6
	succResult := []int{0, 2}
	fmt.Println("test02 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoSum(nums, target))
	fmt.Println()
}

func test03() {
	nums := []int{-1, 0}
	target := -1
	succResult := []int{0, 1}
	fmt.Println("test03 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoSum(nums, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
