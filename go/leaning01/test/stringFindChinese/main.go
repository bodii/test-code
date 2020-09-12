package main

import (
	"fmt"
	"unicode"
)


func main() {
	s := "a23中国人"

	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
