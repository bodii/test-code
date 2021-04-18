package main

import "fmt"

func failure_example() {
	var m map[string]int
	// m := make(map[string]int)
	m["one"] = 1
	fmt.Println(m)
}

func seccess_example() {
	var s []int
	s = append(s, 1)
	fmt.Println(s)
}

func main() {
	seccess_example()
}
