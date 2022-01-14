package main

import (
	"fmt"
)

func hanota(A []int, B []int, C []int) []int {
	if A == nil {
		return nil
	}

	var move func(n int, a, b, c *[]int)
	move = func(num int, from, mid, to *[]int) {
		if num == 1 {
			*to, *from = append(*to, (*from)[len(*from)-1]), (*from)[:len(*from)-1]
			return
		}

		move(num-1, from, to, mid)
		move(1, from, mid, to)
		move(num-1, mid, from, to)
	}

	move(len(A), &A, &B, &C)

	return C
}

func test01() {
	A, B, C := []int{2, 1, 0}, []int{}, []int{}
	succResult := []int{2, 1, 0}

	fmt.Println("test01:")
	fmt.Println("A:", A, "B:", B, "C:", C)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", hanota(A, B, C))
	fmt.Println()
}

func test02() {
	A, B, C := []int{1, 0}, []int{}, []int{}
	succResult := []int{1, 0}

	fmt.Println("test02:")
	fmt.Println("A:", A, "B:", B, "C:", C)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", hanota(A, B, C))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
