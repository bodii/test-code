package main

import "fmt"

func failure_example() {
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	if v := x["two"]; v == "" {
		fmt.Println("key two is entry")
	}
}

func success_example() {
	x := map[string]string{"one": 2, "two": "", "three": "3"}
	if _, ok := x["two"]; !ok {
		fmt.Println("key two is no entry")
	}
}

func main() {
	success_example()
}
