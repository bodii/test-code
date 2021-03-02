package main

import (
	"fmt"
	"time"
)

func main() {
	n := 2

	start := time.Now()
	bad := firstBadVersion(n)
	fmt.Println(bad)
	fmt.Println(time.Now().Sub(start))
}

func isBadVersion(n int) bool {
	return n >= 1
}

func firstBadVersion(n int) int {
	left, mid, right := 1, 0, n

	for right > left {
		mid = left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
