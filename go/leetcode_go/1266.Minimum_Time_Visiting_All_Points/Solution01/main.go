package main

import (
	"fmt"
)

func minTimeToVisitAllPoints(points [][]int) int {
	count := 0

	for i := 0; i < len(points)-1; i++ {
		count += max(abs(points[i][0]-points[i+1][0]), abs(points[i][1]-points[i+1][1]))
	}

	return count
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	succResult := 7
	fmt.Println("test01 points:=", points)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minTimeToVisitAllPoints(points))
	fmt.Println()
}

func test02() {
	points := [][]int{{3, 2}, {-2, 2}}
	succResult := 5
	fmt.Println("test02 points:=", points)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minTimeToVisitAllPoints(points))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
