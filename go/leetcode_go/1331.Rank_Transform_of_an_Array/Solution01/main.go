package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	arrLen := len(arr)
	if arrLen == 0 {
		return []int{}
	}

	arrSort := make([]int, arrLen)
	copy(arrSort, arr)
	sort.Ints(arrSort)

	arrSortMap := make(map[int]int)
	index := 1
	for _, v := range arrSort {
		if arrSortMap[v] == 0 {
			arrSortMap[v] = index
			index++
		}

	}

	result := make([]int, arrLen)
	for k, v := range arr {
		result[k] = arrSortMap[v]
	}

	return result
}

func test01() {
	arr := []int{40, 10, 20, 30}
	succResult := []int{4, 1, 2, 3}
	fmt.Println("test01 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arrayRankTransform(arr))
	fmt.Println()
}

func test02() {
	arr := []int{100, 100, 100}
	succResult := []int{1, 1, 1}
	fmt.Println("test02 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arrayRankTransform(arr))
	fmt.Println()
}

func test03() {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	succResult := []int{5, 3, 4, 2, 8, 6, 7, 1, 3}
	fmt.Println("test03 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arrayRankTransform(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
