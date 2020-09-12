package main

import (
	"fmt"
)

func main() {
	fmt.Printf("5=%b\n", 5)
	fmt.Printf("2=%b\n", 2)
	fmt.Println(5 & 2)
	fmt.Println(5 | 2)
	fmt.Println(5 ^ 2)
	fmt.Println(5 << 2)
	fmt.Println(5 >> 2)
}
