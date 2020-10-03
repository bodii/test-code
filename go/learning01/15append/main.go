package main

import (
	"fmt"
)

func main() {
	s1 := []string{"beijing", "shanghai"}
	s1 = append(s1, "guangzhou")
	fmt.Println(s1)

	s2 := []string{"hangzhou", "chengdu"}
	s1 = append(s1, s2...)
	fmt.Println(s1, len(s1), cap(s1))
}
