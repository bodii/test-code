package main

import (
	"fmt"
)

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}

func test01() {
	n := 16
	succResult := true
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfFour(n))
	fmt.Println()
}

func test02() {
	n := 5
	succResult := false
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfFour(n))
	fmt.Println()
}

func test03() {
	n := 1
	succResult := true
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfFour(n))
	fmt.Println()
}

func test04() {
	n := 8
	succResult := false
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfFour(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
