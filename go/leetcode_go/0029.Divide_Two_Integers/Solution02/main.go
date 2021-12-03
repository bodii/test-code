package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == divisor {
		return 1
	}

	if dividend == -divisor || -dividend == divisor {
		return -1
	}

	isNegative := false
	if dividend < 0 {
		isNegative = !isNegative
		dividend = -dividend
	}

	if divisor < 0 {
		isNegative = !isNegative
		divisor = -divisor
	}

	v, i := divisor, 1
	result := 0
	for v <= dividend {
		if v+v <= dividend {
			v <<= 1
			i <<= 1
		} else {
			result += i
			dividend -= v
			v, i = divisor, 1
		}
	}

	if isNegative {
		result = -result
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}

func test01() {
	dividend, divisor := 10, 3
	succResult := 3
	fmt.Println("test01 dividend:=", dividend, " divisor:=", divisor)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(dividend, divisor))
	fmt.Println()
}

func test02() {
	dividend, divisor := 7, -3
	succResult := -2
	fmt.Println("test02 dividend:=", dividend, " divisor:=", divisor)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(dividend, divisor))
	fmt.Println()
}

func test03() {
	dividend, divisor := -7, -3
	succResult := 2
	fmt.Println("test03 dividend:=", dividend, " divisor:=", divisor)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(dividend, divisor))
	fmt.Println()
}

func test04() {
	dividend, divisor := 1, 1
	succResult := 1
	fmt.Println("test04 dividend:=", dividend, " divisor:=", divisor)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(dividend, divisor))
	fmt.Println()
}

func test05() {
	dividend, divisor := 2147483647, 1
	succResult := 2147483647
	fmt.Println("test05 dividend:=", dividend, " divisor:=", divisor)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(dividend, divisor))
	fmt.Println()
}

func test06() {
	dividend, divisor := 5, -1
	succResult := -5
	fmt.Println("test06 dividend:=", dividend, " divisor:=", divisor)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(dividend, divisor))
	fmt.Println()
}

func test07() {
	dividend, divisor := -2147483648, -1
	succResult := 2147483647
	fmt.Println("test07 dividend:=", dividend, " divisor:=", divisor)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(dividend, divisor))
	fmt.Println()
}

func test08() {
	dividend, divisor := 6543, 2
	succResult := 3271
	fmt.Println("test08 dividend:=", dividend, " divisor:=", divisor)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", divide(dividend, divisor))
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
