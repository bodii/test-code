package main

import (
	"fmt"
	"math/bits"
)

func isPowerOfFour(n int) bool {
	len := bits.Len32(uint32(n - 1))
	if len&1 == 1 {
		return false
	}

	return n == 1<<len
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

func test05() {
	n := 4
	succResult := true
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfFour(n))
	fmt.Println()
}

func test06() {
	n := 6
	succResult := false
	fmt.Println("test06 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfFour(n))
	fmt.Println()
}

func test07() {
	n := 24
	succResult := false
	fmt.Println("test07 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfFour(n))
	fmt.Println()
}

func test08() {
	n := 0
	succResult := false
	fmt.Println("test08 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfFour(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
}
