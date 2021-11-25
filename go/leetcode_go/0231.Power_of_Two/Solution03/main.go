package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	pow := 0
	for i := 0; ; i++ {
		pow = 1 << i
		if n < pow {
			return false
		} else if n == pow {
			return true
		}
	}
}

func test01() {
	n := 1
	succResult := true
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfTwo(n))
	fmt.Println()
}

func test02() {
	n := 16
	succResult := true
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfTwo(n))
	fmt.Println()
}

func test03() {
	n := 3
	succResult := false
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfTwo(n))
	fmt.Println()
}

func test04() {
	n := 4
	succResult := true
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfTwo(n))
	fmt.Println()
}

func test05() {
	n := 5
	succResult := false
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPowerOfTwo(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
