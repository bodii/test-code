package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v >= target {
			return k
		}
	}

	return len(nums)
}

func test01() {
	nums := []int{1, 3, 5, 6}
	target := 5
	succResult := 2
	fmt.Println("test01 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", searchInsert(nums, target))
	fmt.Println()
}

func test02() {
	nums := []int{1, 3, 5, 6}
	target := 2
	succResult := 1
	fmt.Println("test02 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", searchInsert(nums, target))
	fmt.Println()
}

func test03() {
	nums := []int{1, 3, 5, 6}
	target := 7
	succResult := 4
	fmt.Println("test03 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", searchInsert(nums, target))
	fmt.Println()
}

func test04() {
	nums := []int{1, 3, 5, 6}
	target := 0
	succResult := 0
	fmt.Println("test04 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", searchInsert(nums, target))
	fmt.Println()
}

func test05() {
	nums := []int{1}
	target := 0
	succResult := 0
	fmt.Println("test05 nums:=", nums, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", searchInsert(nums, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
