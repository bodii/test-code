package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	numsMap := make(map[int]int)
	half := len(nums) / 2
	for _, v := range nums {
		numsMap[v]++
		if numsMap[v] > half {
			return v
		}
	}

	return -1
}

func test01() {
	nums := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	succResult := 5
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", majorityElement(nums))
	fmt.Println()
}

func test02() {
	nums := []int{3, 2}
	succResult := -1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", majorityElement(nums))
	fmt.Println()
}

func test03() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	succResult := 2
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", majorityElement(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
