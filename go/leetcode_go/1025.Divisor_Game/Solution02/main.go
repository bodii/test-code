package main

import (
	"fmt"
)

func divisorGame(n int) bool {
	d := make([]bool, n+5)

	d[1] = false
	d[2] = true
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !d[i-j] {
				d[i] = true
				break
			}
		}
	}

	return d[n]
}

func test01() {
	n := 2
	successResult := true
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", divisorGame(n))
	fmt.Println()
}

func test02() {
	n := 3
	successResult := false
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", divisorGame(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
