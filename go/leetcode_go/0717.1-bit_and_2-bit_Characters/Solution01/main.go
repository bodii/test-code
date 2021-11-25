package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	end := len(bits) - 1
	i := 0
	for i < end {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}

	return i == end
}

func test01() {
	bits := []int{1, 0, 0}
	succResult := true
	fmt.Println("test01 bits:=", bits)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isOneBitCharacter(bits))
	fmt.Println()
}

func test02() {
	bits := []int{1, 1, 1, 0}
	succResult := false
	fmt.Println("test02 bits:=", bits)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isOneBitCharacter(bits))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
