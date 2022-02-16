package main

import (
	"fmt"
	"math"
)

func divide(a int, b int) int {
	if a == 0 {
		return 0
	}

	isNegative := false
	if a < 0 {
		a = -a
		isNegative = !isNegative
	}

	if b < 0 {
		b = -b
		isNegative = !isNegative
	}

	result := 0
	for i := 31; i >= 0; i-- {
		if b<<i <= a {
			a -= b << i
			result += 1 << i
		}
	}

	if isNegative {
		result = -result
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		result = math.MaxInt32
	}

	return result
}

func test01() {
	a, b := 15, 2
	succResult := 7
	fmt.Println("test01 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(a, b))
	fmt.Println()
}

func test02() {
	a, b := 7, -3
	succResult := -2
	fmt.Println("test02 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(a, b))
	fmt.Println()
}

func test03() {
	a, b := 0, 1
	succResult := 0
	fmt.Println("test03 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(a, b))
	fmt.Println()
}

func test04() {
	a, b := 1, 1
	succResult := 1
	fmt.Println("test04 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(a, b))
	fmt.Println()
}

func test05() {
	a, b := 2147483647, 1
	succResult := 2147483647
	fmt.Println("test05 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(a, b))
	fmt.Println()
}

func test06() {
	a, b := 5, -1
	succResult := -5
	fmt.Println("test06 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(a, b))
	fmt.Println()
}

func test07() {
	a, b := -2147483648, -1
	succResult := 2147483647
	fmt.Println("test07 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(a, b))
	fmt.Println()
}

func test08() {
	a, b := 6543, 2
	succResult := 3271
	fmt.Println("test08 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(a, b))
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
