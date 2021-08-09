package main

import (
	"fmt"
)

func maximum(a int, b int) int {
	c := (a - b) >> 63
	fmt.Println((3 - 1) >> 63)
	return a*(c+1) - b*c
}

func test01() {
	a, b := 1, 2
	succResult := 2
	fmt.Println("test01 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximum(a, b))
	fmt.Println()
}

func main() {
	test01()
}
