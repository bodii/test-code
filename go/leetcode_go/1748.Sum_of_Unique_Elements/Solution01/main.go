package main

import (
	"fmt"
)

func sumOfUnique(nums []int) int {
	numsMap := make(map[int]int)

	for _, v := range nums {
		numsMap[v]++
	}

	res := 0
	for k, v := range numsMap {
		if v == 1 {
			res += k
		}
	}

	return res
}

func test01() {
	nums := []int{1, 2, 3, 2}
	succResult := 4
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumOfUnique(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 1, 1, 1}
	succResult := 0
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumOfUnique(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 3, 4, 5}
	succResult := 15
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumOfUnique(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
