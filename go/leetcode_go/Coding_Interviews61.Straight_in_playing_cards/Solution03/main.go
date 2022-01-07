package main

import (
	"fmt"
)

func isStraight(nums []int) bool {
	numsMap := make([]int, 14)

	for _, n := range nums {
		numsMap[n]++

		if n != 0 && numsMap[n] > 1 {
			return false
		}
	}

	l, r := 1, 13
	for numsMap[l] == 0 {
		l++
	}

	for numsMap[r] == 0 {
		r--
	}

	return r-l < 5
}

func test01() {
	nums := []int{1, 2, 3, 4, 5}
	succResult := true
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isStraight(nums))
	fmt.Println()
}

func test02() {
	nums := []int{0, 0, 1, 2, 5}
	succResult := true
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isStraight(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
