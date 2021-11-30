package main

import (
	"fmt"
)

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1

	for l < r {
		mid := l + (r-l)/2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}

	return numbers[l]
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
