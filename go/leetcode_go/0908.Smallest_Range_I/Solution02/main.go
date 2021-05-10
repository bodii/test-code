package main

import (
	"fmt"
	"sort"
)

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)

	diff := A[len(A)-1] - A[0] - K*2
	if diff <= 0 {
		return 0
	}

	return diff
}

func test01() {
	A := []int{1}
	K := 0
	success := 0
	succResult := []int{1}

	fmt.Println("test01 A:=", A, " K:=", K)
	fmt.Println("success:=", success)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallestRangeI(A, K))
	fmt.Println()
}

func test02() {
	A := []int{0, 10}
	K := 2
	success := 6
	succResult := []int{2, 8}

	fmt.Println("test02 A:=", A, " K:=", K)
	fmt.Println("success:=", success)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallestRangeI(A, K))
	fmt.Println()
}

func test03() {
	A := []int{1, 3, 6}
	K := 3
	success := 0
	succResult := [][]int{{3, 3, 3}, {4, 4, 4}}

	fmt.Println("test03 A:=", A, " K:=", K)
	fmt.Println("success:=", success)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", smallestRangeI(A, K))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
