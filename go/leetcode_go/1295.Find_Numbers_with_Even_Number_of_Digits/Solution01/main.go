package main

import (
	"fmt"
)

func findNumbers(nums []int) int {
	result := 0
	for _, num := range nums {
		len := findLen(num)
		if len > 1 && len%2 == 0 {
			result++
		}
	}

	return result
}

func findLen(num int) int {
	if 1 <= num && num < 10 {
		return 1
	} else if 10 <= num && num < 100 {
		return 2
	} else if 100 <= num && num < 1000 {
		return 3
	} else if 1000 <= num && num < 10000 {
		return 4
	} else if 10000 <= num && num < 100000 {
		return 5
	} else if 100000 <= num {
		return 6
	}
	return 0
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
