package main

import (
	"fmt"
)

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}

	return true
}

func test01() {
	n := 6
	succResult := true
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUgly(n))
	fmt.Println()
}

func test02() {
	n := 8
	succResult := true
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUgly(n))
	fmt.Println()
}

func test03() {
	n := 14
	succResult := false
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUgly(n))
	fmt.Println()
}

func test04() {
	n := 1
	succResult := true
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUgly(n))
	fmt.Println()
}

func test05() {
	n := 0
	succResult := false
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUgly(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
