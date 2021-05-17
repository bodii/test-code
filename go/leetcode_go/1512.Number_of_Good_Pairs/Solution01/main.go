package main

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	numsLen := len(nums)
	result := 0

	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if i < j && nums[i] == nums[j] {
				result++
			}
		}
	}

	return result
}

func test01() {
	nums := []int{1, 2, 3, 1, 1, 3}
	succResult := 4
	fmt.Println("test01 arr:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numIdenticalPairs(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 1, 1}
	succResult := 6
	fmt.Println("test02 arr:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numIdenticalPairs(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 3}
	succResult := 0
	fmt.Println("test03 arr:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numIdenticalPairs(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
