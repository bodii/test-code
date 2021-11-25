package main

import (
	"fmt"
	"math/bits"
)

func convertInteger(A int, B int) int {
	return bits.OnesCount32(uint32(A ^ B))
}

func test01() {
	a, b := 29, 15
	succResult := 2
	fmt.Println("test01 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", convertInteger(a, b))
	fmt.Println()
}

func test02() {
	a, b := 1, 2
	succResult := 2
	fmt.Println("test02 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", convertInteger(a, b))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
