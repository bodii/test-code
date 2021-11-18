package main

import (
	"fmt"
)

func countNegatives(grid [][]int) int {
	colIdx := len(grid[0]) - 1
	num := 0
	for rowIdx := len(grid) - 1; rowIdx >= 0 && grid[rowIdx][colIdx] < 0; rowIdx-- {
		for j := colIdx; j >= 0 && grid[rowIdx][j] < 0; j-- {
			num++
		}
	}

	return num
}

func test01() {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	succResult := 8
	fmt.Println("test01 grid:=", grid)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countNegatives(grid))
	fmt.Println()
}

func test02() {
	grid := [][]int{{3, 2}, {1, 0}}
	succResult := 0
	fmt.Println("test02 grid:=", grid)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countNegatives(grid))
	fmt.Println()
}

func test03() {
	grid := [][]int{{1, -1}, {-1, -1}}
	succResult := 3
	fmt.Println("test03 grid:=", grid)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countNegatives(grid))
	fmt.Println()
}

func test04() {
	grid := [][]int{{-1}}
	succResult := 1
	fmt.Println("test04 grid:=", grid)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countNegatives(grid))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
