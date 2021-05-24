package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	result := 0
	for _, num := range nums {
		len := len(strconv.Itoa(num))
		if len > 1 && len%2 == 0 {
			result++
		}
	}

	return result
}

func test01() {
	nums := []int{12, 345, 2, 6, 7896}
	succResult := 2
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findNumbers(nums))
	fmt.Println()
}

func test02() {
	nums := []int{555, 901, 482, 1771}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findNumbers(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
