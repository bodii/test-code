package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool)
	for _, v := range nums {
		if numsMap[v] {
			return true
		}
		numsMap[v] = true
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
