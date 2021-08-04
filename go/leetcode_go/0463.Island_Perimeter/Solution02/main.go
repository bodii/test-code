package main

import (
	"fmt"
)

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

/* 深度优先搜索 */
func islandPerimeter(grid [][]int) int {
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				return dfs(i, j, &grid)
			}
		}
	}

	return 0
}

func dfs(i, j int, grid *[][]int) int {
	// 超出边界
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) {
		return 1
	}

	// 一条边
	if (*grid)[i][j] == 0 {
		return 1
	}

	// 已访问过
	if (*grid)[i][j] == 2 {
		return 0
	}
	// 标记访问过
	(*grid)[i][j] = 2

	return dfs(i-1, j, grid) +
		dfs(i+1, j, grid) +
		dfs(i, j-1, grid) +
		dfs(i, j+1, grid)
}

func test01() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	succResult := 16
	fmt.Println("test01 grid:=", grid)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", islandPerimeter(grid))
	fmt.Println()
}

func test02() {
	grid := [][]int{{1}}
	succResult := 4
	fmt.Println("test02 grid:=", grid)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", islandPerimeter(grid))
	fmt.Println()
}

func test03() {
	grid := [][]int{{1, 0}}
	succResult := 4
	fmt.Println("test03 grid:=", grid)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", islandPerimeter(grid))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
