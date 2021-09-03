package main

import (
	"fmt"
)

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceSum := getSliceIntsSum(aliceSizes)
	bobSum := getSliceIntsSum(bobSizes)

	aliceMap := make(map[int]bool)
	for _, v := range aliceSizes {
		if !aliceMap[v] {
			aliceMap[v] = true
		}
	}

	d := (bobSum - aliceSum) / 2
	for _, b := range bobSizes {
		if aliceMap[b-d] {
			return []int{b - d, b}
		}
	}

	return []int{0, 0}
}

func getSliceIntsSum(s []int) int {
	sum := 0

	for _, v := range s {
		sum += v
	}

	return sum
}

func test01() {
	aliceSizes := []int{1, 1}
	bobSizes := []int{2, 2}
	succResult := []int{1, 2}
	fmt.Println("test01 aliceSizes:=", aliceSizes, " bobSizes:=", bobSizes)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fairCandySwap(aliceSizes, bobSizes))
	fmt.Println()
}

func test02() {
	aliceSizes := []int{1, 2}
	bobSizes := []int{2, 3}
	succResult := []int{1, 2}
	fmt.Println("test02 aliceSizes:=", aliceSizes, " bobSizes:=", bobSizes)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fairCandySwap(aliceSizes, bobSizes))
	fmt.Println()
}

func test03() {
	aliceSizes := []int{2}
	bobSizes := []int{1, 3}
	succResult := []int{2, 3}
	fmt.Println("test03 aliceSizes:=", aliceSizes, " bobSizes:=", bobSizes)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fairCandySwap(aliceSizes, bobSizes))
	fmt.Println()
}

func test04() {
	aliceSizes := []int{1, 2, 5}
	bobSizes := []int{2, 4}
	succResult := []int{5, 4}
	fmt.Println("test04 aliceSizes:=", aliceSizes, " bobSizes:=", bobSizes)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fairCandySwap(aliceSizes, bobSizes))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
