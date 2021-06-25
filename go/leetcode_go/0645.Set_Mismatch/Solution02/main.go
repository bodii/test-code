package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	numsLen := len(nums)
	nums2 := make([]int, numsLen)
	for i := 0; i < numsLen; i++ {
		nums2[nums[i]-1]++
	}

	for k, v := range nums2 {
		if v == 0 {
			result[1] = k + 1
		} else if v == 2 {
			result[0] = k + 1
		}
	}

	return result
}

func test01() {
	nums := []int{1, 2, 2, 4}
	succResult := []int{2, 3}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1}
	succResult := []int{1, 2}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 2, 3, 5, 4}
	succResult := []int{2, 6}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func test04() {
	nums := []int{2, 2}
	succResult := []int{2, 1}
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func test05() {
	nums := []int{2, 2, 3}
	succResult := []int{2, 1}
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func test06() {
	nums := []int{3, 2, 3, 4, 6, 5}
	succResult := []int{3, 1}
	fmt.Println("test06 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findErrorNums(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
