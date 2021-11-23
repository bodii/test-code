package main

import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	soldiers := make([]int, len(mat))
	indexs := make([]int, len(mat))
	for i, m := range mat {
		for _, v := range m {
			soldiers[i] += v
		}
		indexs[i] = i
	}

	sort.SliceStable(indexs, func(i, j int) bool {
		return soldiers[indexs[i]] < soldiers[indexs[j]] ||
			(soldiers[indexs[i]] == soldiers[indexs[j]] && indexs[i] < indexs[j])
	})

	return indexs[:k]
}

func test01() {
	k := 3
	mat := [][]int{{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}
	succResult := []int{2, 0, 3}
	fmt.Println("test01 k:=", k, " mat:=", mat)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kWeakestRows(mat, k))
	fmt.Println()
}

func test02() {
	k := 2
	mat := [][]int{{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0}}
	succResult := []int{0, 2}
	fmt.Println("test02 k:=", k, " mat:=", mat)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kWeakestRows(mat, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
