package main

import (
	"container/list"
	"fmt"
)

func processQueries(queries []int, m int) []int {
	nodeList := list.New()
	result := make([]int, len(queries))

	for i := 1; i <= m; i++ {
		nodeList.PushBack(i)
	}

	getIndex := func(val int) int {
		index := 0
		for e := nodeList.Front(); e != nil; e = e.Next() {
			if e.Value == val {
				nodeList.MoveToFront(e)
				break
			}

			index++
		}

		return index
	}

	for k, v := range queries {
		result[k] = getIndex(v)
	}

	return result
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
