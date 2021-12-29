package main

import (
	"fmt"
)

func countPoints(rings string) int {
	counts := make([]int, 10)
	colors := map[byte]int{'R': 1, 'G': 2, 'B': 4}
	for i := 0; i < len(rings); i += 2 {
		counts[rings[i+1]-'0'] |= colors[rings[i]]
	}

	nums := 0
	for _, c := range counts {
		if c == 7 {
			nums++
		}
	}

	return nums
}

func test01() {
	s := "B0B6G0R6R0R6G9"
	succResult := 1
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countPoints(s))
	fmt.Println()
}

func test02() {
	s := "B0R0G0R9R0B0G0"
	succResult := 1
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countPoints(s))
	fmt.Println()
}

func test03() {
	s := "G4"
	succResult := 0
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countPoints(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
