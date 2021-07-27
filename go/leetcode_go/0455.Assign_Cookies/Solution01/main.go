package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(g)))
	sort.Ints(s)

	result, j := 0, len(s)-1
	for i := 0; i < len(g); i++ {
		if j < 0 {
			return result
		}
		if s[j] >= g[i] {
			result++
			j--
		}
	}

	return result
}

func test01() {
	g, s := []int{1, 2, 3}, []int{1, 1}
	succResult := 1
	fmt.Println("test01 g:=", g, " s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findContentChildren(g, s))
	fmt.Println()
}

func test02() {
	g, s := []int{1, 2}, []int{1, 2, 3}
	succResult := 2
	fmt.Println("test02 g:=", g, " s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findContentChildren(g, s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
