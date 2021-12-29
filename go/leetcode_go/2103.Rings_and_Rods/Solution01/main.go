package main

import (
	"fmt"
)

type rgb [3]int

func countPoints(rings string) int {
	counts := make([]rgb, 10)
	colors := map[byte]int{'R': 0, 'G': 1, 'B': 2}
	for i := 0; i < len(rings); i += 2 {
		counts[rings[i+1]-'0'][colors[rings[i]]]++
	}

	nums := 0
	for _, c := range counts {
		if c[0] > 0 && c[1] > 0 && c[2] > 0 {
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
