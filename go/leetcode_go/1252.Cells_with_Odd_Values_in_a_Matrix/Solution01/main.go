package main

import (
	"fmt"
)

func oddCells(m int, n int, indices [][]int) int {
	rec := make([][]int, m)
	for i := 0; i < m; i++ {
		rec[i] = make([]int, n)
	}

	for _, ind := range indices {
		for i := 0; i < n; i++ {
			rec[ind[0]][i]++
		}
		for i := 0; i < m; i++ {
			rec[i][ind[1]]++
		}
	}

	oddNums := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rec[i][j]&1 == 1 {
				oddNums++
			}
		}
	}

	return oddNums
}

func test01() {
	m, n := 2, 3
	indices := [][]int{{0, 1}, {1, 1}}
	succResult := 6
	fmt.Println("test01 indices:=", indices)
	fmt.Println("m:=", m, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", oddCells(m, n, indices))
	fmt.Println()
}

func test02() {
	m, n := 2, 2
	indices := [][]int{{1, 1}, {0, 0}}
	succResult := 0
	fmt.Println("test02 indices:=", indices)
	fmt.Println("m:=", m, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", oddCells(m, n, indices))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
