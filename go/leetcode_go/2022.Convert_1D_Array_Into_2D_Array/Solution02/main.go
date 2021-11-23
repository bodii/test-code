package main

import (
	"fmt"
)

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = original[n*i : n*(i+1)]
	}

	return result
}

func test01() {
	original := []int{1, 2, 3, 4}
	m, n := 2, 2
	succResult := [][]int{{1, 2}, {3, 4}}
	fmt.Println("test01 original:=", original)
	fmt.Println("m:=", m, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", construct2DArray(original, m, n))
	fmt.Println()
}

func test02() {
	original := []int{1, 2, 3}
	m, n := 1, 3
	succResult := [][]int{{1, 2, 3}}
	fmt.Println("test02 original:=", original)
	fmt.Println("m:=", m, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", construct2DArray(original, m, n))
	fmt.Println()
}

func test03() {
	original := []int{1, 2}
	m, n := 1, 1
	succResult := [][]int{}
	fmt.Println("test03 original:=", original)
	fmt.Println("m:=", m, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", construct2DArray(original, m, n))
	fmt.Println()
}

func test04() {
	original := []int{3}
	m, n := 1, 2
	succResult := [][]int{}
	fmt.Println("test04 original:=", original)
	fmt.Println("m:=", m, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", construct2DArray(original, m, n))
	fmt.Println()
}

func test05() {
	original := []int{1, 1, 1, 1}
	m, n := 4, 1
	succResult := [][]int{{1}, {1}, {1}, {1}}
	fmt.Println("test05 original:=", original)
	fmt.Println("m:=", m, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", construct2DArray(original, m, n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
