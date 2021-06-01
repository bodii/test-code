package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arrMap := make(map[int]int)
	for _, v := range arr1 {
		arrMap[v]++
	}

	resultArr := make([]int, 0)
	for _, v := range arr2 {
		for i := 0; i < arrMap[v]; i++ {
			resultArr = append(resultArr, v)
		}
		delete(arrMap, v)
	}

	tailArr := make([]int, 0)
	for k, v := range arrMap {
		for i := 0; i < v; i++ {
			tailArr = append(tailArr, k)
		}
	}

	sort.Ints(tailArr)

	resultArr = append(resultArr, tailArr...)

	return resultArr
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
