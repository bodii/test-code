package main

import "fmt"

func numMovesStones(a int, b int, c int) []int {
	min, mid, max := a, b, c

	if min > mid {
		min, mid = mid, min
	}

	if mid > max {
		mid, max = max, mid
	}

	if min > mid {
		min, mid = mid, min
	}

	left, right := mid-min-1, max-mid-1
	if left == 0 && right == 0 {
		return []int{0, 0}
	}

	maxMove := left + right
	if left <= 1 || right <= 1 {
		return []int{1, maxMove}
	}

	return []int{2, maxMove}
}

func test01() {
	a, b, c := 1, 2, 5
	succResult := []int{1, 2}

	fmt.Println("test01 a:=", a, " b:=", b, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numMovesStones(a, b, c))
	fmt.Println()
}

func test02() {
	a, b, c := 4, 3, 2
	succResult := []int{0, 0}

	fmt.Println("test02 a:=", a, " b:=", b, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numMovesStones(a, b, c))
	fmt.Println()
}

func test03() {
	a, b, c := 3, 5, 1
	succResult := []int{1, 2}

	fmt.Println("test03 a:=", a, " b:=", b, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numMovesStones(a, b, c))
	fmt.Println()
}

func test04() {
	a, b, c := 7, 4, 1
	succResult := []int{2, 4}

	fmt.Println("test04 a:=", a, " b:=", b, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numMovesStones(a, b, c))
	fmt.Println()
}

func test05() {
	a, b, c := 1, 3, 6
	succResult := []int{1, 3}

	fmt.Println("test05 a:=", a, " b:=", b, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numMovesStones(a, b, c))
	fmt.Println()
}

func test06() {
	a, b, c := 1, 4, 6
	succResult := []int{1, 3}

	fmt.Println("test06 a:=", a, " b:=", b, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numMovesStones(a, b, c))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
