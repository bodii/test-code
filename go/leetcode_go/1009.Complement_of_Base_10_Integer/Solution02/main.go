package main

import (
	"fmt"
)

func bitwiseComplement(n int) int {
	num := 1
	// n 二进制=101， num二进制＝111, n^m = 10 = 2(十进制)
	for num < n {
		num <<= 1
		num += 1
	}
	return num ^ n
}

func test01() {
	n := 5
	succResult := 2
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", bitwiseComplement(n))
	fmt.Println()
}

func test02() {
	n := 7
	succResult := 0
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", bitwiseComplement(n))
	fmt.Println()
}

func test03() {
	n := 10
	succResult := 5
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", bitwiseComplement(n))
	fmt.Println()
}

func test04() {
	n := 0
	succResult := 1
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", bitwiseComplement(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
