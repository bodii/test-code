package main

import (
	"fmt"
	"sort"
)

func luckyNumbers(matrix [][]int) []int {
	mins, maxs := make([]int, len(matrix)), make([]int, len(matrix[0]))

	for i, m := range matrix {
		minVal := matrix[i][0]
		for j, v := range m {
			mins[i], maxs[j] = min(minVal, v), max(maxs[j], v)
			minVal = mins[i]
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(mins)))
	sort.Ints(maxs)
	if mins[0] == maxs[0] {
		return mins[:1]
	}

	return []int{}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	matrix := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	succResult := []int{15}
	fmt.Println("test01 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func test02() {
	matrix := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	succResult := []int{12}
	fmt.Println("test02 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func test03() {
	matrix := [][]int{{7, 8}, {1, 2}}
	succResult := []int{7}
	fmt.Println("test03 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func test04() {
	matrix := [][]int{{3, 6}, {7, 1}, {5, 2}, {4, 8}}
	succResult := []int{}
	fmt.Println("test04 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func test05() {
	matrix := [][]int{{36376, 85652, 21002, 4510}, {68246, 64237, 42962, 9974}, {32768, 97721, 47338, 5841}, {55103, 18179, 79062, 46542}}
	succResult := []int{}
	fmt.Println("test05 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
