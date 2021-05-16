package main

import (
	"fmt"
)

func xorOperation(n int, start int) int {
	result := start

	for i := 1; i < n; i++ {
		result ^= (start + 2*i)
	}

	return result
}

func test01() {
	n, start := 5, 0
	succResult := 8
	fmt.Println("test01 n:=", n, " start:=", start)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", xorOperation(n, start))
	fmt.Println()
}

func test02() {
	n, start := 4, 3
	succResult := 8
	fmt.Println("test02 n:=", n, " start:=", start)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", xorOperation(n, start))
	fmt.Println()
}

func test03() {
	n, start := 1, 7
	succResult := 7
	fmt.Println("test03 n:=", n, " start:=", start)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", xorOperation(n, start))
	fmt.Println()
}

func test04() {
	n, start := 10, 5
	succResult := 2
	fmt.Println("test04 n:=", n, " start:=", start)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", xorOperation(n, start))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
