package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr1Map := make(map[int]int)
	for _, v := range arr1 {
		arr1Map[v]++
	}

	result := make([]int, len(arr1))
	j := 0
	for _, v := range arr2 {
		if arr1Map[v] > 0 {
			for i := 0; i < arr1Map[v]; i++ {
				result[j] = v
				j++
			}
		}
		arr1Map[v] = 0
	}
	result = result[0:j]

	noIns := make([]int, 0)
	for v, n := range arr1Map {
		if n == 0 {
			continue
		}

		for i := 0; i < n; i++ {
			noIns = append(noIns, v)
		}
	}

	sort.Ints(noIns)
	return append(result, noIns...)
}

func test01() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	succResult := []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}
	fmt.Println("test01 arr1:=", arr1)
	fmt.Println("test01 arr2:=", arr2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", relativeSortArray(arr1, arr2))
	fmt.Println()
}

func main() {
	test01()
}
