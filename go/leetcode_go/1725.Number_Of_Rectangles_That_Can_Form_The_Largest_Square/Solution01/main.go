package main

import (
	"fmt"
)

func countGoodRectangles(rectangles [][]int) int {
	lens := make(map[int]int)
	for _, rect := range rectangles {
		lens[min(rect[0], rect[1])]++
	}

	maxLen, maxNum := 0, 0
	for len := range lens {
		maxLen = max(maxLen, len)
		maxNum = lens[maxLen]
	}

	return maxNum
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
	rectangles := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	succResult := 3
	fmt.Println("test01 rectangles:=", rectangles)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countGoodRectangles(rectangles))
	fmt.Println()
}

func test02() {
	rectangles := [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}
	succResult := 3
	fmt.Println("test02 rectangles:=", rectangles)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countGoodRectangles(rectangles))
	fmt.Println()
}

func test03() {
	rectangles := [][]int{{5, 8}, {3, 9}, {3, 12}}
	succResult := 1
	fmt.Println("test03 rectangles:=", rectangles)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countGoodRectangles(rectangles))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
