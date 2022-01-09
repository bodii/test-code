package main

import (
	"fmt"
)

func isArmstrong(n int) bool {
	sum := 0
	for n2 := n; n2 > 0; n2 /= 10 {
		val := n2 % 10

		sum += val * val * val
	}

	return sum == n
}

func test01() {
	n := 153
	succResult := true
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isArmstrong(n))
	fmt.Println()
}

func test02() {
	n := 123
	succResult := false
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isArmstrong(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
