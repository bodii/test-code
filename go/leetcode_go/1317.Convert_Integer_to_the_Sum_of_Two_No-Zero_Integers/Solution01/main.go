package main

import (
	"fmt"
)

func getNoZeroIntegers(n int) []int {
	for i := 1; i <= n; i++ {
		if isNoZero(i) && isNoZero(n-i) {
			return []int{i, n - i}
		}
	}

	return []int{}
}

func isNoZero(n int) bool {
	if n == 0 {
		return false
	}

	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n /= 10
	}

	return true
}

func test01() {
	n := 2
	succResult := []int{1, 1}
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getNoZeroIntegers(n))
	fmt.Println()
}

func test02() {
	n := 11
	succResult := []int{2, 9}
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getNoZeroIntegers(n))
	fmt.Println()
}

func test03() {
	n := 10000
	succResult := []int{1, 9999}
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getNoZeroIntegers(n))
	fmt.Println()
}

func test04() {
	n := 69
	succResult := []int{1, 68}
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getNoZeroIntegers(n))
	fmt.Println()
}

func test05() {
	n := 1010
	succResult := []int{11, 999}
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getNoZeroIntegers(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
