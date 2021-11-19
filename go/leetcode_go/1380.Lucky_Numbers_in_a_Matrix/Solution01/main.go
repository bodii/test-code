package main

import (
	"fmt"
)

func luckyNumbers(matrix [][]int) []int {
	for _, m := range matrix {
		minNum, minIndex := findMinNumAndIndex(m)

		maxNum := matrix[0][minIndex]
		for j := 1; j < len(matrix); j++ {
			if maxNum < matrix[j][minIndex] {
				maxNum = matrix[j][minIndex]
			}
		}

		if maxNum == minNum {
			return []int{maxNum}
		}
	}

	return []int{}
}

func findMinNumAndIndex(list []int) (min, index int) {
	min = list[0]
	for i := 1; i < len(list); i++ {
		if min > list[i] {
			min, index = list[i], i
		}
	}

	return
}

func test01() {
	matrix := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	succResult := []int{15}
	fmt.Println("test01 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func test02() {
	matrix := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	succResult := []int{12}
	fmt.Println("test02 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func test03() {
	matrix := [][]int{{7, 8}, {1, 2}}
	succResult := []int{7}
	fmt.Println("test03 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func test04() {
	matrix := [][]int{{3, 6}, {7, 1}, {5, 2}, {4, 8}}
	succResult := []int{}
	fmt.Println("test04 matrix:=", matrix)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", luckyNumbers(matrix))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
