package main

import (
	"fmt"
)

func main() {
	var m1 = make(map[string][]string, 3)
	m1["k1"] = []string{"val1"}
	m1["k1"] = append(m1["k1"], "val1")
	m1["k1"] = append(m1["k1"], "val2")

	fmt.Println(m1)
}
