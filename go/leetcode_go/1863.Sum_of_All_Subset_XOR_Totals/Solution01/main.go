package main

import (
	"fmt"
)

func subsetXORSum(nums []int) int {
	sum := 0
	numsLen := len(nums)

	for _, v := range nums {
		sum |= v
	}

	return sum << (numsLen - 1)
}

func test01() {
	nums := []int{1, 3}
	succResult := 6
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", subsetXORSum(nums))
	fmt.Println()
}

func test02() {
	nums := []int{5, 1, 6}
	succResult := 28
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", subsetXORSum(nums))
	fmt.Println()
}

func test03() {
	nums := []int{3, 4, 5, 6, 7, 8}
	succResult := 480
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", subsetXORSum(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
