package main

import (
	"fmt"
)

func totalMoney(n int) int {
	result := 0

	weeks, lastWeekMaxDay := n/7, n%7
	maxDay := 7
	for i := 0; i <= weeks; i++ {
		if i == weeks {
			maxDay = lastWeekMaxDay
		}
		for d := 1; d <= maxDay; d++ {
			result += i + d
		}
	}

	return result
}

func test01() {
	n := 4
	succResult := 10
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", totalMoney(n))
	fmt.Println()
}

func test02() {
	n := 10
	succResult := 37
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", totalMoney(n))
	fmt.Println()
}

func test03() {
	n := 20
	succResult := 96
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", totalMoney(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
