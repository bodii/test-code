package main

import (
	"fmt"
	"sort"
)

func merge(A []int, m int, B []int, n int) {
	copy(A[m:], B[:n])
	sort.Ints(A)

	fmt.Println(A)
}

func test01() {
	A, m := []int{1, 2, 3, 0, 0, 0}, 3
	B, n := []int{2, 5, 6}, 3
	succResult := []int{1, 2, 2, 3, 5, 6}
	fmt.Println("test01 A:=", A, " B:=", B)
	fmt.Println("test01 m:=", m, " n:=", n)
	fmt.Println("success result:=", succResult)
	merge(A, m, B, n)

	fmt.Println()
}

func main() {
	test01()
}
