package main

import (
	"fmt"
)

func merge(A []int, m int, B []int, n int) {
	l, r := 0, 0
	list := make([]int, m+n)
	for l < m || r < n {
		if l == m || A[l] > B[r] {
			list[r+l] = B[r]
			r++
		} else if r == n || A[l] <= B[r] {
			list[r+l] = A[l]
			l++
		}
	}

	copy(A, list)

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
