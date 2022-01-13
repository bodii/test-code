package main

import (
	"fmt"
	"math/bits"
)

func countPrimeSetBits(left int, right int) int {
	primes := []int{0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1}

	count := 0
	for ; left <= right; left++ {
		if primes[bits.OnesCount(uint(left))] == 1 {
			count++
		}
	}

	return count
}

func test01() {
	left, right := 6, 10
	succResult := 4
	fmt.Println("test01:")
	fmt.Println("left:", left, " right:=", right)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", countPrimeSetBits(left, right))
	fmt.Println()
}

func test02() {
	left, right := 10, 15
	succResult := 5
	fmt.Println("test02:")
	fmt.Println("left:", left, " right:=", right)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", countPrimeSetBits(left, right))
	fmt.Println()
}

func test03() {
	left, right := 842, 888
	succResult := 23
	fmt.Println("test03:")
	fmt.Println("left:", left, " right:=", right)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", countPrimeSetBits(left, right))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
