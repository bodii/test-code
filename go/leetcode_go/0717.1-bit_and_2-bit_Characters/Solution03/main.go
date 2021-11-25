package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	last := len(bits)
	for i := last - 2; i >= 0; i-- {
		if bits[i] == 0 {
			return (last-i)&1 == 0
		}
	}
	return (last-1)&1 == 0
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
