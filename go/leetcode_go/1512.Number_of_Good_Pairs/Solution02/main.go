package main

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	result := 0
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}

	for _, v := range numsMap {
		if v < 2 {
			continue
		}

		result += (v - 1) * v / 2
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
func test04() {
	nums := []int{1, 1, 1, 1, 1, 1}
	succResult := 15
	fmt.Println("test04 arr:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numIdenticalPairs(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
