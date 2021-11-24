package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	result := 0

	for i, v := range tickets {
		if i > k {
			result += min(tickets[k]-1, v)
		} else {
			result += min(tickets[k], v)
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func test01() {
	tickets := []int{2, 3, 2}
	k := 2
	succResult := 6

	fmt.Println("test01 tickets:=", tickets, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", timeRequiredToBuy(tickets, k))
	fmt.Println()
}

func test02() {
	tickets := []int{5, 1, 1, 1}
	k := 0
	succResult := 8

	fmt.Println("test02 tickets:=", tickets, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", timeRequiredToBuy(tickets, k))
	fmt.Println()
}

func test03() {
	tickets := []int{84, 49, 5, 24, 70, 77, 87, 8}
	k := 3
	succResult := 154

	fmt.Println("test03 tickets:=", tickets, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", timeRequiredToBuy(tickets, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
