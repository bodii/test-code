package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	result := deck[len(deck)-1:]
	for len(deck) > 1 {
		deck = deck[:len(deck)-1]
		result = append(deck[len(deck)-1:], append(result[len(result)-1:], result[:len(result)-1]...)...)
	}

	return result
}

func test01() {
	deck := []int{17, 13, 11, 2, 3, 5, 7}
	success := []int{2, 13, 3, 11, 5, 17, 7}

	fmt.Println("test01  deck:=", deck)
	fmt.Println("success result:", success)
	fmt.Println("result:", deckRevealedIncreasing(deck))
	fmt.Println()
}

func main() {
	test01()
}
