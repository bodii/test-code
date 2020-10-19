package main

import "fmt"

func trailingZeroes(n int) int {
	zeroes := 0
	for n > 0 {
		n /= 5
		zeroes += n
	}
	return zeroes
}

func main() {
	fmt.Println(trailingZeroes(3))
}
