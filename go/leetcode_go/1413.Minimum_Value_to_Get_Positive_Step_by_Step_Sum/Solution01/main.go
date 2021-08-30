package main

import (
	"fmt"
)

func minStartValue(nums []int) int {
	sum, min := 0, 1

	for _, v := range nums {
		sum += v
		if sum < min {
			min = sum
		}
	}

	if min < 1 {
		return 1 - min
	}

	return min
}

func test01() {
	nums := []int{-3, 2, -3, 4, 2}
	succResult := 5
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minStartValue(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minStartValue(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, -2, -3}
	succResult := 5
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minStartValue(nums))
	fmt.Println()
}

func test04() {
	nums := []int{2, 3, 5, -5, -1}
	succResult := 1
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minStartValue(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
