package main

import (
	"fmt"
	"math"
)

func getMinDistance(nums []int, target int, start int) int {
	min := math.MaxInt64

	for k, v := range nums {
		if v == target {
			d := abs(k - start)
			if d < min {
				min = d
			}
		}
	}

	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func test01() {
	nums := []int{1, 2, 3, 4, 5}
	target, start := 5, 3
	succResult := 1
	fmt.Println("test01 nums:=", nums, " targe:=", target, " start:=", start)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getMinDistance(nums, target, start))
	fmt.Println()
}

func test02() {
	nums := []int{1}
	target, start := 1, 0
	succResult := 0
	fmt.Println("test02 nums:=", nums, " targe:=", target, " start:=", start)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getMinDistance(nums, target, start))
	fmt.Println()
}

func test03() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target, start := 1, 0
	succResult := 0
	fmt.Println("test03 nums:=", nums, " targe:=", target, " start:=", start)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getMinDistance(nums, target, start))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
