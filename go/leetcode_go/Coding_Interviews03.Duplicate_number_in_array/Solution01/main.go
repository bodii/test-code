package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	numsMpa := make(map[int]bool)
	for _, n := range nums {
		if numsMpa[n] {
			return n
		}

		numsMpa[n] = true
	}

	return -1
}

func test01() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	succResult := "2 或 3"
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findRepeatNumber(nums))
	fmt.Println()
}

func main() {
	test01()
}
