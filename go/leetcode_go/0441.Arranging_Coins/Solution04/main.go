package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	return int(math.Sqrt(float64(8*n+1))-1) / 2
}

func test01() {
	n := 5
	successResult := 2
	fmt.Println("test01 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", arrangeCoins(n))
	fmt.Println()
}

func test02() {
	n := 8
	successResult := 3
	fmt.Println("test02 n:", n)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", arrangeCoins(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
