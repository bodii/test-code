package main

import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	dict := make(map[int]int)
	for _, v := range arr {
		dict[v]++
	}

	frequencys := make([]int, 0)
	for _, v := range dict {
		frequencys = append(frequencys, v)
	}

	sort.Ints(frequencys)

	for i := 0; i < len(frequencys)-1; i++ {
		if frequencys[i] == frequencys[i+1] {
			return false
		}
	}

	return true
}

func test01() {
	arr := []int{1, 2, 2, 1, 1, 3}
	success := true

	fmt.Println("test01  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", uniqueOccurrences(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1, 2}
	success := false

	fmt.Println("test02  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", uniqueOccurrences(arr))
	fmt.Println()
}

func test03() {
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	success := true

	fmt.Println("test03  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", uniqueOccurrences(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
