package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	prev := -1
	for n > 0 {
		if n&1 == prev {
			return false
		}
		prev = n & 1
		n >>= 1
	}

	return true
}

func test01() {
	n := 5
	succResult := true
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hasAlternatingBits(n))
	fmt.Println()
}

func test02() {
	n := 7
	succResult := false
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hasAlternatingBits(n))
	fmt.Println()
}

func test03() {
	n := 11
	succResult := false
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hasAlternatingBits(n))
	fmt.Println()
}

func test04() {
	n := 10
	succResult := true
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hasAlternatingBits(n))
	fmt.Println()
}

func test05() {
	n := 3
	succResult := false
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hasAlternatingBits(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
