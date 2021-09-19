package main

import (
	"fmt"
)

func add(a int, b int) int {
	if b == 0 {
		return a
	}
	return add(a^b, (a&b)<<1)
}

func test01() {
	a, b := 1, 1
	succResult := 2
	fmt.Println("test01 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", add(a, b))
	fmt.Println()
}

func test02() {
	a, b := 10, 17
	succResult := 27
	fmt.Println("test02 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", add(a, b))
	fmt.Println()
}

func test03() {
	a, b := 88, 10
	succResult := 98
	fmt.Println("test03 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", add(a, b))
	fmt.Println()
}

func test04() {
	a, b := -1, 2
	succResult := 1
	fmt.Println("test04 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", add(a, b))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
