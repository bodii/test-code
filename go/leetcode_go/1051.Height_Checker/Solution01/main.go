package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	heightsLen := len(heights)
	swapNums := 0

	heightsCopy := make([]int, heightsLen)
	copy(heightsCopy, heights)
	sort.Ints(heightsCopy)

	for i := 0; i < heightsLen; i++ {
		if heights[i] != heightsCopy[i] {
			swapNums++
		}
	}

	return swapNums
}

func test01() {
	heights := []int{1, 1, 4, 2, 1, 3}
	successResult := 3
	fmt.Println("test01 heights:", heights)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", heightChecker(heights))
	fmt.Println()
}

func test02() {
	heights := []int{5, 1, 2, 3, 4}
	successResult := 5
	fmt.Println("test02 heights:", heights)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", heightChecker(heights))
	fmt.Println()
}

func test03() {
	heights := []int{1, 2, 3, 4, 5}
	successResult := 0
	fmt.Println("test03 heights:", heights)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", heightChecker(heights))
	fmt.Println()
}

func test04() {
	heights := []int{2, 1, 2, 1, 1, 2, 2, 1}
	successResult := 4
	fmt.Println("test04 heights:", heights)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", heightChecker(heights))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
