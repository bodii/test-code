package main

import (
	"fmt"
	"sort"
)

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceSum := getSliceIntsSum(aliceSizes)
	bobSum := getSliceIntsSum(bobSizes)

	sort.Ints(aliceSizes)
	sort.Ints(bobSizes)

	i, j := 0, 0
	for i < len(aliceSizes) && j < len(bobSizes) {
		if aliceSum-aliceSizes[i]+bobSizes[j] < bobSum-bobSizes[j]+aliceSizes[i] {
			j++
		} else if aliceSum-aliceSizes[i]+bobSizes[j] > bobSum-bobSizes[j]+aliceSizes[i] {
			i++
		} else {
			return []int{aliceSizes[i], bobSizes[j]}
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
