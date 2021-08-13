package main

import (
	"fmt"
)

func findMagicIndex(nums []int) int {
	for k, v := range nums {
		if k == v {
			return v
		}
	}

	return -1
}

func test01() {
	nums := []int{0, 2, 3, 4, 5}
	succResult := 0
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMagicIndex(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 1}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMagicIndex(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
