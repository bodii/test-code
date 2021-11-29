package main

import (
	"fmt"
)

func matrixReshape(mat [][]int, r int, c int) [][]int {
	area := len(mat) * len(mat[0])
	if area != r*c {
		return mat
	}

	result := make([][]int, r)
	for k := range result {
		result[k] = make([]int, c)
	}

	m := len(mat[0])
	for i := 0; i < area; i++ {
		result[i/c][i%c] = mat[i/m][i%m]
	}

	return result
}

func test01() {
	mat := [][]int{{1, 2}, {3, 4}}
	r, c := 1, 4
	succResult := [][]int{{1, 2, 3, 4}}
	fmt.Println("test01 mat:=", mat, " r:=", r, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", matrixReshape(mat, r, c))
	fmt.Println()
}

func test02() {
	mat := [][]int{{1, 2}, {3, 4}}
	r, c := 2, 4
	succResult := [][]int{{1, 2}, {3, 4}}
	fmt.Println("test02 mat:=", mat, " r:=", r, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", matrixReshape(mat, r, c))
	fmt.Println()
}

func test03() {
	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}
	r, c := 42, 5
	succResult := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}
	fmt.Println("test03 mat:=", mat, " r:=", r, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", matrixReshape(mat, r, c))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
