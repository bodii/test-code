package main

import (
	"fmt"
	"sort"
)

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	result := make([][]int, rows*cols)

	for i, k := 0, 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[k] = []int{i, j}
			k++
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return abs(result[i][0]-rCenter)+abs(result[i][1]-cCenter) < abs(result[j][0]-rCenter)+abs(result[j][1]-cCenter)
	})

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func test01() {
	rows, cols := 1, 2
	rCenter, cCenter := 0, 0
	succResult := [][]int{{0, 0}, {0, 1}}
	fmt.Println("test01:")
	fmt.Printf("rows:=%d, cols:=%d, rCenter:=%d,cCenter:=%d\n", rows, cols, rCenter, cCenter)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", allCellsDistOrder(rows, cols, rCenter, cCenter))
	fmt.Println()
}

func test02() {
	rows, cols := 2, 2
	rCenter, cCenter := 0, 1
	succResult := [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}
	fmt.Println("test02:")
	fmt.Printf("rows:=%d, cols:=%d, rCenter:=%d,cCenter:=%d\n", rows, cols, rCenter, cCenter)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", allCellsDistOrder(rows, cols, rCenter, cCenter))
	fmt.Println()
}

func test03() {
	rows, cols := 2, 3
	rCenter, cCenter := 1, 2
	succResult := [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}
	fmt.Println("test03:")
	fmt.Printf("rows:=%d, cols:=%d, rCenter:=%d,cCenter:=%d\n", rows, cols, rCenter, cCenter)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", allCellsDistOrder(rows, cols, rCenter, cCenter))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
