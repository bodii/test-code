package main

import (
	"fmt"
)

func decode(encoded []int, first int) []int {
	encodedLen := len(encoded)

	result := make([]int, encodedLen+1)
	result[0] = first

	for i := 0; i < encodedLen; i++ {
		result[i+1] = encoded[i] ^ result[i]
	}

	return result
}

func test01() {
	encoded := []int{1, 2, 3}
	first := 1
	succResult := []int{1, 0, 2, 1}
	fmt.Println("test01 encoded:=", encoded, " first:=", first)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", decode(encoded, first))
	fmt.Println()
}

func test02() {
	encoded := []int{6, 2, 7, 3}
	first := 4
	succResult := []int{4, 2, 0, 7, 4}
	fmt.Println("test02 encoded:=", encoded, " first:=", first)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", decode(encoded, first))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
