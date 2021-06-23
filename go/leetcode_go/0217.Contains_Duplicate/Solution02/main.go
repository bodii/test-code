package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
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

func main() {
	test01()
	test02()
	test03()
}
