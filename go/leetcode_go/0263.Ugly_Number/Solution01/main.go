package main

import (
	"fmt"
)

func isUgly(n int) bool {
	for n > 0 {
		n1 := n / 2
		if n1*2 == n {
			n /= 2
			continue
		}
		n2 := n / 3
		if n2*3 == n {
			n /= 3
			continue
		}
		n3 := n / 5
		if n3*5 == n {
			n /= 5
			continue
		}

		if n == 1 {
			return true
		}

		return false
	}

	return false
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
