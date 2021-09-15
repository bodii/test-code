package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	numMap := make(map[int]int)
	numsLen := len(nums)
	for i := 0; i < numsLen; i += 2 {
		if i < numsLen {
			numMap[nums[i]]++
		}
		if i+1 < numsLen {
			numMap[nums[i+1]]++
		}
	}

	return numMap[target]
}

func test01() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	succResult := 2
	fmt.Println("test01 target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", search(nums, target))
	fmt.Println()
}

func test02() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	succResult := 0
	fmt.Println("test02 target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", search(nums, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
