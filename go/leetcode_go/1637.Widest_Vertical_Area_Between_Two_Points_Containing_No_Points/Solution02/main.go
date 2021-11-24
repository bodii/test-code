package main

import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	len := len(points)
	widths := make([]int, len)
	for k, v := range points {
		widths[k] = v[0]
	}
	sort.Ints(widths)

	result := 0
	for i := 0; i < len-1; i++ {
		result = max(result, widths[i+1]-widths[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	points := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	succResult := 1
	fmt.Println("test01 points:=", points)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxWidthOfVerticalArea(points))
	fmt.Println()
}

func test02() {
	points := [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}
	succResult := 3
	fmt.Println("test02 points:=", points)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxWidthOfVerticalArea(points))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
