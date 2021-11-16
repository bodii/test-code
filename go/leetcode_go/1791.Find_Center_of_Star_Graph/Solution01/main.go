package main

import (
	"fmt"
)

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][1] == edges[1][0] {
		return edges[1][0]
	}

	return edges[1][1]
}

func test01() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	succResult := 2
	fmt.Println("test01 edges:=", edges)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findCenter(edges))
	fmt.Println()
}

func test02() {
	edges := [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}
	succResult := 1
	fmt.Println("test02 edges:=", edges)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findCenter(edges))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
