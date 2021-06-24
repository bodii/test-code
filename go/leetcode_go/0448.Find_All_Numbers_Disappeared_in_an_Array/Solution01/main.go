package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	numsMap := make(map[int]bool)
	for _, v := range nums {
		numsMap[v] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !numsMap[i] {
			result = append(result, i)
		}
	}

	return result
}

func test01() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	succResult := []int{5, 6}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findDisappearedNumbers(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1}
	succResult := []int{2}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findDisappearedNumbers(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
