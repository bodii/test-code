package main

import (
	"fmt"
)

func isHappy(n int) bool {
	for {
		fmt.Println(n)
		if n < 10 {
			if n == 1 || n == 7 {
				return true
			} else {
				return false
			}
		} else {
			v := n
			v2 := 0
			for v > 0 {
				v3 := v % 10
				v2 += v3 * v3
				v /= 10
			}

			n = v2
		}
	}
}

func test01() {
	n := 19
	succResult := true
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isHappy(n))
	fmt.Println()
}

func test02() {
	n := 2
	succResult := false
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isHappy(n))
	fmt.Println()
}

func test03() {
	n := 4
	succResult := false
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isHappy(n))
	fmt.Println()
}

func test04() {
	n := 7
	succResult := true
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isHappy(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
