package main

import (
	"fmt"
)

func rotatedDigits(N int) int {
	isRotated := func(t int) bool {
		flag := false
		for t > 0 {
			switch t % 10 {
			case 3, 4, 7:
				return false
			case 2, 5, 6, 9:
				flag = true
			}
			t /= 10
		}
		return flag
	}

	sum := 0
	for i := 1; i <= N; i++ {
		if isRotated(i) {
			sum++
		}
	}

	return sum
}

func test01() {
	N := 10
	succResult := 4
	fmt.Println("test01 N:=", N)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", rotatedDigits(N))
	fmt.Println()
}

func test02() {
	N := 2
	succResult := 1
	fmt.Println("test02 N:=", N)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", rotatedDigits(N))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
