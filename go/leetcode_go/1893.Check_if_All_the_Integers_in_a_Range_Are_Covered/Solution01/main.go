package main

import (
	"fmt"
)

func isCovered(ranges [][]int, left int, right int) bool {
	list := make([]int, 52)
	for _, vals := range ranges {
		list[vals[0]]++
		list[vals[1]+1]--
	}

	sum := 0
	for i := 0; i <= 50; i++ {
		sum += list[i]
		if sum <= 0 && left <= i && right >= i {
			return false
		}
	}

	return true
}

func test01() {
	ranges := [][]int{{1, 2}, {3, 4}, {5, 6}}
	left, right := 2, 5
	succResult := true
	fmt.Println("test01 ranges:=", ranges, " left:=", left, " right:=", right)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isCovered(ranges, left, right))
	fmt.Println()
}

func test02() {
	ranges := [][]int{{1, 10}, {10, 20}}
	left, right := 21, 21
	succResult := false
	fmt.Println("test02 ranges:=", ranges, " left:=", left, " right:=", right)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isCovered(ranges, left, right))
	fmt.Println()
}

func test03() {
	ranges := [][]int{{1, 1}}
	left, right := 1, 50
	succResult := false
	fmt.Println("test03 ranges:=", ranges, " left:=", left, " right:=", right)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isCovered(ranges, left, right))
	fmt.Println()
}

func test04() {
	ranges := [][]int{{1, 50}}
	left, right := 1, 50
	succResult := true
	fmt.Println("test04 ranges:=", ranges, " left:=", left, " right:=", right)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isCovered(ranges, left, right))
	fmt.Println()
}

func test05() {
	ranges := [][]int{{25, 42}, {7, 14}, {2, 32}, {25, 28}, {39, 49}, {1, 50}, {29, 45}, {18, 47}}
	left, right := 15, 38
	succResult := true
	fmt.Println("test05 ranges:=", ranges, " left:=", left, " right:=", right)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isCovered(ranges, left, right))
	fmt.Println()
}

func test06() {
	ranges := [][]int{{37, 49}, {5, 17}, {8, 32}}
	left, right := 29, 49
	succResult := false
	fmt.Println("test06 ranges:=", ranges, " left:=", left, " right:=", right)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isCovered(ranges, left, right))
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
