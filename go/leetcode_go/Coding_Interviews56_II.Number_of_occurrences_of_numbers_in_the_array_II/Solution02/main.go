package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}

	for k, v := range numsMap {
		if v == 1 {
			return k
		}
	}

	return -1
}

func test01() {
	nums := []int{3, 4, 3, 3}
	succResult := 4
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", singleNumber(nums))
	fmt.Println()
}

func test02() {
	nums := []int{9, 1, 7, 9, 7, 9, 7}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", singleNumber(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
