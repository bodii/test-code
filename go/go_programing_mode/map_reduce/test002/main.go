package main

import (
	"fmt"
)

// Reduce ...
func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += len(it)
	}

	return sum
}

func main() {
	list := []string{"Jack", "Tome", "Lias"}

	z := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", z)
}
