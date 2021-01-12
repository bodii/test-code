package main

import "fmt"

// Filter ...
func Filter(arr []int, fn func(n int) bool) []int {
	newArray := []int{}

	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}

	return newArray
}

func main() {
	intSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	out := Filter(intSet, func(n int) bool {
		return n%2 == 1
	})

	fmt.Printf("%v\n", out)

	out = Filter(intSet, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n", out)
}
