package main

import (
	"fmt"
	"strings"
)


func main() {
	s := "abcddcba"
	s1 := strings.Split(s, "")
	b := true

	half := len(s) / 2
	j := len(s) - 1
	for i:=0;i<half;i++ {
		if s1[i] != s1[j] {
			b = false
			fmt.Println(s1[i], s1[j])
		}
		j --
	}
	fmt.Println(b)


	b1 := true
	for k := range s {
		if s[k] != s[len(s) - 1 - k] {
			b1 = false
			fmt.Println(s[k], s[len(s) - 1 - k])
		}
	}
	fmt.Println(b1)
}
