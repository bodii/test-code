package main

import (
	"fmt"
)

func distributeCandies(candyType []int) int {
	mid := len(candyType) / 2
	types := make(map[int]int)

	for _, v := range candyType {
		types[v] = 1
	}

	if mid <= len(types) {
		return mid
	}

	return len(types)
}

func test01() {
	candyType := []int{1, 1, 2, 2, 3, 3}
	succResult := 3
	fmt.Println("test01 candyType:=", candyType)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", distributeCandies(candyType))
	fmt.Println()
}

func test02() {
	candyType := []int{1, 1, 2, 3}
	succResult := 2
	fmt.Println("test02 candyType:=", candyType)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", distributeCandies(candyType))
	fmt.Println()
}

func test03() {
	candyType := []int{6, 6, 6, 6}
	succResult := 1
	fmt.Println("test03 candyType:=", candyType)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", distributeCandies(candyType))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
