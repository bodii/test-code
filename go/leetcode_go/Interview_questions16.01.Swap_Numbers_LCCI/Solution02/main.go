package main

import (
	"fmt"
)

func swapNumbers(numbers []int) []int {
	if numbers[0] == numbers[1] {
		return numbers
	}

	numbers[1] += numbers[0]
	numbers[0] = numbers[1] - numbers[0]
	numbers[1] = numbers[1] - numbers[0]

	return numbers
}

func test01() {
	nums := []int{1, 2}
	succResult := []int{2, 1}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", swapNumbers(nums))
	fmt.Println()
}

func main() {
	test01()
}
