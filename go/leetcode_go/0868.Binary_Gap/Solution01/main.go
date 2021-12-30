package main

import (
	"fmt"
)

func binaryGap(n int) int {
	b := fmt.Sprintf("%b", n)

	maxGap, start := 0, 0
	for i := 1; i < len(b); i++ {
		if b[i] == '1' {
			maxGap = max(maxGap, i-start)
			start = i
		}
	}

	return maxGap
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	n := 22
	successResult := 2
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", binaryGap(n))
	fmt.Println()
}

func test02() {
	n := 5
	successResult := 2
	fmt.Println("test02 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", binaryGap(n))
	fmt.Println()
}

func test03() {
	n := 6
	successResult := 1
	fmt.Println("test03 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", binaryGap(n))
	fmt.Println()
}

func test04() {
	n := 8
	successResult := 0
	fmt.Println("test04 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", binaryGap(n))
	fmt.Println()
}

func test05() {
	n := 1
	successResult := 0
	fmt.Println("test05 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", binaryGap(n))
	fmt.Println()
}

func test06() {
	n := 3281
	successResult := 4
	fmt.Println("test06 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", binaryGap(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
