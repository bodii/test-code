package main

import (
	"fmt"
)

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	minRow, minCol := m, n
	for _, v := range ops {
		minRow = min(minRow, v[0])
		minCol = min(minCol, v[1])
	}

	return minRow * minCol
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func test01() {
	m, n := 3, 3
	ops := [][]int{{2, 2}, {3, 3}}
	succResult := 4
	fmt.Println("test01 ops:=", ops)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxCount(m, n, ops))
	fmt.Println()
}

func test02() {
	m, n := 3, 3
	ops := [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}
	succResult := 4
	fmt.Println("test02 ops:=", ops)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxCount(m, n, ops))
	fmt.Println()
}

func test03() {
	m, n := 3, 3
	ops := [][]int{}
	succResult := 9
	fmt.Println("test03 ops:=", ops)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxCount(m, n, ops))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
