package main

import "fmt"

func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for i := 1; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		} else if A[i] > max {
			max = A[i]
		}
	}

	diff := max - min - K*2
	if 0 >= diff {
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
