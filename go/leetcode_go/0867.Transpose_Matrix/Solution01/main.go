package main

import (
	"fmt"
)

func transpose(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[j] = append(result[j], matrix[i][j])
		}
	}

	return result
}

func test01() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	succResult := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	fmt.Println("test01 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", transpose(matrix))
	fmt.Println()
}

func test02() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	succResult := [][]int{{1, 4}, {2, 5}, {3, 6}}
	fmt.Println("test02 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", transpose(matrix))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
