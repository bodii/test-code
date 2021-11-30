package main

import (
	"fmt"
)

func minArray(numbers []int) int {
	r := len(numbers)

	for i := 1; i < r; i++ {
		if numbers[i-1] > numbers[i] {
			return numbers[i]
		}
	}

	return numbers[0]
}

func test01() {
	nums := []int{3, 4, 5, 1, 2}
	succResult := 1
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minArray(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 2, 2, 0, 1}
	succResult := 0
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minArray(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
