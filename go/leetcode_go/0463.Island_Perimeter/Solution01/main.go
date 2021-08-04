package main

import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
	result := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}

			result += 4

			if i > 0 && grid[i-1][j] == 1 {
				result--
			}

			if i+1 < m && grid[i+1][j] == 1 {
				result--
			}

			if j > 0 && grid[i][j-1] == 1 {
				result--
			}

			if j+1 < n && grid[i][j+1] == 1 {
				result--
			}
		}
	}

	return result
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
