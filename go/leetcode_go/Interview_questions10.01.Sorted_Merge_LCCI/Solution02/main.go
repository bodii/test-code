package main

import (
	"fmt"
)

func merge(A []int, m int, B []int, n int) {
	maxIndex := m + n - 1
	for i, j := 0, 0; i <= maxIndex && j < n; i++ {
		if A[i] > B[j] {
			for k := maxIndex; k > i; k-- {
				A[k] = A[k-1]
			}
			A[i] = B[j]
			j++
			m++
		} else if i >= m {
			A[i] = B[j]
			j++
		}
	}

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
