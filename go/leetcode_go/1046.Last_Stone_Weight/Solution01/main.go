package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	lastIdx := len(stones) - 1
	if lastIdx == 0 {
		return stones[lastIdx]
	}
	sort.Ints(stones)

	for i := lastIdx - 1; i > 0; i-- {
		stones[i+1], stones[i] = 0, stones[i+1]-stones[i]
		if stones[i-1] > stones[i] {
			sort.Ints(stones)
		}

		fmt.Println(stones)
	}

	return stones[0]
}

func test01() {
	stones := []int{2, 7, 4, 1, 8, 1}
	succResult := 1
	fmt.Println("test01 stones:=", stones)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lastStoneWeight(stones))
	fmt.Println()
}

func test02() {
	stones := []int{7, 6, 7, 6, 9}
	succResult := 3
	fmt.Println("test02 stones:=", stones)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lastStoneWeight(stones))
	fmt.Println()
}

func test03() {
	stones := []int{9, 10, 1, 7, 3}
	succResult := 2
	fmt.Println("test03 stones:=", stones)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lastStoneWeight(stones))
	fmt.Println()
}

func test04() {
	stones := []int{434, 667, 378, 919, 212, 902, 240, 257, 208, 996, 411, 222, 557, 634, 425, 949, 755, 833, 785, 886, 40, 159, 932, 157, 764, 916, 85, 300, 130, 278}
	succResult := 1
	fmt.Println("test04 stones:=", stones)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lastStoneWeight(stones))
	fmt.Println()
}

func main() {
	// test01()
	// test02()
	// test03()
	test04()
}
