package main

import (
	"fmt"
)

func diagonalSum(mat [][]int) int {
	lineLastIdx := len(mat[0]) - 1

	sum := 0
	for k, m := range mat {
		if k != lineLastIdx-k {
			sum += m[k] + m[lineLastIdx-k]
		} else {
			sum += m[k]
		}
	}

	return sum
}

func test01() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	succResult := 25
	fmt.Println("test01 mat:=", mat)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", diagonalSum(mat))
	fmt.Println()
}

func test02() {
	mat := [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	succResult := 8
	fmt.Println("test02 mat:=", mat)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", diagonalSum(mat))
	fmt.Println()
}

func test03() {
	mat := [][]int{{5}}
	succResult := 5
	fmt.Println("test03 mat:=", mat)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", diagonalSum(mat))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
