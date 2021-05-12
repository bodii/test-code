package main

import "fmt"

func calculate(s string) int {
	x, y := 1, 0

	for _, v := range s {
		if v == 65 {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}

	return x + y
}

func test01() {
	s := "AB"
	succResult := 4
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", calculate(s))
	fmt.Println()
}

func main() {
	test01()
}
