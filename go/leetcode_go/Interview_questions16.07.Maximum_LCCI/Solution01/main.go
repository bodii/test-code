package main

import (
	"fmt"
	"math"
)

func maximum(a int, b int) int {
	return (int(math.Abs(float64(a-b))) + a + b) / 2
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
