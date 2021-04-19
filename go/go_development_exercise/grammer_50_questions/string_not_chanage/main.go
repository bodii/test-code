package main

import "fmt"

func failure_example() {
	x := "text"
	x[0] = "T"
	fmt.Println(x)
}

func success_example() {
	x := "text"
	xBytes := []byte(x)
	xBytes[0] = 'T'
	x = string(xBytes)
	fmt.Println(x)
}

func main() {
	success_example()
}
