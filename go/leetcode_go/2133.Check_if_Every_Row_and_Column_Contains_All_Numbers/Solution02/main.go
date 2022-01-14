package main

import "fmt"

func checkValid(matrix [][]int) bool {
	max := len(matrix)
	if max == 1 {
		return matrix[0][0] == 1
	}

	for i := 0; i < max; i++ {
		rows := make([]bool, max+1)
		cols := make([]bool, max+1)
		for j := 0; j < max; j++ {
			if rows[matrix[i][j]] || cols[matrix[j][i]] {
				return false
			}
			rows[matrix[i][j]], cols[matrix[j][i]] = true, true
		}
	}

	return true
}

func test01() {
	matrix := [][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}
	successResult := true
	fmt.Println("test01 matrix:=", matrix)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", checkValid(matrix))
	fmt.Println()
}

func test02() {
	matrix := [][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}}
	successResult := false
	fmt.Println("test02 matrix:=", matrix)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", checkValid(matrix))
	fmt.Println()
}

func test03() {
	matrix := [][]int{{1}}
	successResult := true
	fmt.Println("test03 matrix:=", matrix)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", checkValid(matrix))
	fmt.Println()
}

func test04() {
	matrix := [][]int{{1, 3, 1}, {3, 1, 3}, {1, 3, 1}}
	successResult := false
	fmt.Println("test04 matrix:=", matrix)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", checkValid(matrix))
	fmt.Println()
}

func test05() {
	matrix := [][]int{{1, 3, 3, 3, 5}, {3, 3, 5, 1, 3}, {5, 3, 3, 3, 1}, {3, 1, 3, 5, 3}, {3, 5, 1, 3, 3}}
	successResult := false
	fmt.Println("test05 matrix:=", matrix)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", checkValid(matrix))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
