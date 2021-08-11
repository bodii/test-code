package main

import (
	"fmt"
)

func expectNumber(scores []int) int {
	resMap := make(map[int]bool)

	for _, v := range scores {
		if !resMap[v] {
			resMap[v] = true
		}
	}

	return len(resMap)
}

func test01() {
	scores := []int{1, 2, 3}
	succResult := 3
	fmt.Println("test01 scores:=", scores)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", expectNumber(scores))
	fmt.Println()
}

func test02() {
	scores := []int{1, 1}
	succResult := 1
	fmt.Println("test02 scores:=", scores)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", expectNumber(scores))
	fmt.Println()
}

func test03() {
	scores := []int{1, 1, 2}
	succResult := 2
	fmt.Println("test03 scores:=", scores)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", expectNumber(scores))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
