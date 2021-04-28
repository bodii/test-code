package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	rows := make([][]int, 0)
	rows = append(rows, []int{1})
	for i := 1; i < numRows; i++ {
		row := make([]int, 0)
		row = append(row, 1)
		for j := 1; j < i; j++ {
			row = append(row, rows[i-1][j-1]+rows[i-1][j])
		}
		row = append(row, 1)
		rows = append(rows, row)
	}

	return rows
}

func test01() {
	fmt.Println("test01 result:")
	numRows := 5
	fmt.Println(generate(numRows))
}

func main() {
	test01()
}
