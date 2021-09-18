package main

import (
	"fmt"
	"sort"
)

func findGCD(nums []int) int {
	sort.Ints(nums)
	min, max := nums[0], nums[len(nums)-1]

	result := 0
	for i := 1; i <= max; i++ {
		if min%i == 0 && max%i == 0 && i > result {
			result = i
		}
	}

	return result
}

func test01() {
	nums := []int{2, 5, 6, 9, 10}
	succResult := 2
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findGCD(nums))
	fmt.Println()
}

func test02() {
	nums := []int{7, 5, 6, 8, 3}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findGCD(nums))
	fmt.Println()
}

func test03() {
	nums := []int{3, 3}
	succResult := 3
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findGCD(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
