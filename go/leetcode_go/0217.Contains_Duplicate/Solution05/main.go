package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func test01() {
	nums := []int{1, 2, 3, 1}
	succResult := true
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsDuplicate(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 3, 4}
	succResult := false
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsDuplicate(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	succResult := true
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsDuplicate(nums))
	fmt.Println()
}

func test04() {
	nums := []int{-100, 1, 7, 3, 9, 4, 6, -100, 5, 2}
	succResult := true
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsDuplicate(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
