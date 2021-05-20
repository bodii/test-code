package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	max1, max2 := 1, 1

	for _, v := range nums {
		if v > max1 {
			max1, max2 = v, max1
		} else if v > max2 {
			max2 = v
		}
	}

	return (max1 - 1) * (max2 - 1)
}

func test01() {
	nums := []int{3, 4, 5, 2}
	succResult := 12
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProduct(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 5, 4, 5}
	succResult := 16
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProduct(nums))
	fmt.Println()
}

func test03() {
	nums := []int{3, 7}
	succResult := 12
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProduct(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
