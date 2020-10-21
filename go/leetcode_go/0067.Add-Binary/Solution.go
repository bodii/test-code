package main

import (
	"fmt"
	"strconv"
)

func addBinary(a, b string) string {
	alen, blen := len(a), len(b)
	n := max(alen, blen)
	result := ""
	tmp := 0
	for i := 0; i < n; i++ {
		if i < alen {
			tmp += int(a[alen-i-1] - '0')
		}
		if i < blen {
			tmp += int(b[blen-i-1] - '0')
		}
		result = strconv.Itoa(tmp%2) + result
		fmt.Println(result)
		tmp /= 2
	}

	if tmp > 0 {
		result = "1" + result
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%v\n", addBinary("100100", "110010"))
}
