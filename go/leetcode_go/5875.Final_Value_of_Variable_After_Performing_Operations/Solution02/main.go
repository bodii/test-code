package main

import (
	"fmt"
)

func finalValueAfterOperations(operations []string) int {
	sum := 0

	for _, v := range operations {
		if v[1] == '+' {
			sum++
		} else {
			sum--
		}
	}

	return sum
}

func test01() {
	operations := []string{"--X", "X++", "X++"}
	succResult := 1
	fmt.Println("test01 operations:=", operations)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", finalValueAfterOperations(operations))
	fmt.Println()
}

func test02() {
	operations := []string{"++X", "++X", "X++"}
	succResult := 3
	fmt.Println("test02 operations:=", operations)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", finalValueAfterOperations(operations))
	fmt.Println()
}

func test03() {
	operations := []string{"X++", "++X", "--X", "X--"}
	succResult := 0
	fmt.Println("test03 operations:=", operations)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", finalValueAfterOperations(operations))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
