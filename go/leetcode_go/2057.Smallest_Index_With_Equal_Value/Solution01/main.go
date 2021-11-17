package main

import (
	"fmt"
)

func smallestEqual(nums []int) int {
	for i, n := range nums {
		if n == i%10 {
			return i
		}
	}

	return -1
}

func test01() {
	nums := []int{0, 1, 2}
	succResult := 0
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallestEqual(nums))
	fmt.Println()
}

func test02() {
	nums := []int{4, 3, 2, 1}
	succResult := 2
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallestEqual(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	succResult := -1
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallestEqual(nums))
	fmt.Println()
}

func test04() {
	nums := []int{2, 1, 3, 5, 2}
	succResult := 1
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallestEqual(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
