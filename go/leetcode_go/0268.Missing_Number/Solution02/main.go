package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	numsLen := len(nums)
	sortNums := make([]int, numsLen+1)

	for _, v := range nums {
		sortNums[v]++
	}

	for k, v := range sortNums {
		if v == 0 {
			return k
		}
	}

	return -1
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
	nums := []int{0, 1}
	succResult := 2
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", missingNumber(nums))
	fmt.Println()
}

func test03() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	succResult := 8
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", missingNumber(nums))
	fmt.Println()
}

func test04() {
	nums := []int{0}
	succResult := 1
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", missingNumber(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
