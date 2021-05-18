package main

import (
	"fmt"
)

func sumBase(n int, k int) int {
	reslut := 0
	for n > 0 {
		reslut += n % k
		n /= k
	}

	return reslut
}

func test01() {
	n, k := 34, 6
	succResult := 9
	fmt.Println("test01 n:=", n, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumBase(n, k))
	fmt.Println()
}

func test02() {
	n, k := 10, 10
	succResult := 1
	fmt.Println("test02 n:=", n, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumBase(n, k))
	fmt.Println()
}

func test03() {
	n, k := 3, 2
	succResult := 2
	fmt.Println("test03 n:=", n, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sumBase(n, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
