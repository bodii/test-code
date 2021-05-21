package main

import (
	"fmt"
)

func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles >= numExchange {
		return numBottles + 1 + (numBottles-numExchange)/(numExchange-1)
	}

	return numBottles
}

func test01() {
	numBottles, numExchange := 9, 3
	succResult := 13
	fmt.Println("test01 numBottles:=", numBottles, " numExchange:=", numExchange)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numWaterBottles(numBottles, numExchange))
	fmt.Println()
}

func test02() {
	numBottles, numExchange := 15, 4
	succResult := 19
	fmt.Println("test02 numBottles:=", numBottles, " numExchange:=", numExchange)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numWaterBottles(numBottles, numExchange))
	fmt.Println()
}

func test03() {
	numBottles, numExchange := 5, 5
	succResult := 6
	fmt.Println("test03 numBottles:=", numBottles, " numExchange:=", numExchange)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numWaterBottles(numBottles, numExchange))
	fmt.Println()
}

func test04() {
	numBottles, numExchange := 2, 3
	succResult := 2
	fmt.Println("test04 numBottles:=", numBottles, " numExchange:=", numExchange)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numWaterBottles(numBottles, numExchange))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
