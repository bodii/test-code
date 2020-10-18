package main

import "fmt"

func titleToNumber(s string) int {
	var result int
	for _, v := range s {
		result = result*26 + (int(v-'A') + 1)
	}
	return result
}

func main() {
	fmt.Println(titleToNumber("ABAZC"))
}
