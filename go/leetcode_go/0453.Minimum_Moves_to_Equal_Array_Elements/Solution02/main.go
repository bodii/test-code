package main

import (
	"fmt"
	"math"
)

func minMoves(nums []int) int {
	sum, min := 0, math.MaxInt32
	for _, v := range nums {
		sum += v
		if v < min {
			min = v
		}
	}

	return sum - min*len(nums)
}

func test01() {
	nums := []int{1, 2, 3}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMoves(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 1000000000}
	succResult := 999999999
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMoves(nums))
	fmt.Println()
}

func test03() {
	nums := []int{-1000000000, -1}
	succResult := 999999999
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMoves(nums))
	fmt.Println()
}

func test04() {
	nums := []int{5, 6, 8, 8, 5}
	succResult := 7
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMoves(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
