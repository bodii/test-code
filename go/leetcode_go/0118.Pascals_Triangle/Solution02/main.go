package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	rows := make([][]int, numRows)
	rows[0] = []int{1}
	for i := 1; i < numRows; i++ {
		rows[i] = make([]int, i+1)
		rows[i][0] = 1
		for j := 1; j < i; j++ {
			rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows[i][i] = 1
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
