package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	numsSlice := make([]int, 100001)
	for _, n := range nums {
		if numsSlice[n] == 1 {
			return n
		}

		numsSlice[n] = 1
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
