package main

import (
	"fmt"
)

func processQueries(queries []int, m int) []int {
	p := make(map[int]int)
	result := make([]int, len(queries))
	for i := 1; i <= m; i++ {
		p[i] = i - 1
	}

	for k, v := range queries {
		result[k] = p[v]
		p = move(p, v)
		fmt.Println(p)
	}

	return result
}

func move(p map[int]int, m int) map[int]int {
	if p[m] == 0 {
		return p
	}

	for k, v := range p {
		if v < p[m] {
			p[k]++
		}
	}

	p[m] = 0

	return p
}

func test01() {
	queries := []int{3, 1, 2, 1}
	m := 5
	succResult := []int{2, 1, 2, 1}
	fmt.Println("test01 queries:=", queries, " m:=", m)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", processQueries(queries, m))
	fmt.Println()
}

func test02() {
	queries := []int{4, 1, 2, 2}
	m := 4
	succResult := []int{3, 1, 2, 0}
	fmt.Println("test02 queries:=", queries, " m:=", m)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", processQueries(queries, m))
	fmt.Println()
}

func test03() {
	queries := []int{7, 5, 5, 8, 3}
	m := 8
	succResult := []int{6, 5, 0, 7, 5}
	fmt.Println("test03 queries:=", queries, " m:=", m)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", processQueries(queries, m))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
