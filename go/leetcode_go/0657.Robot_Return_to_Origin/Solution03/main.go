package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	up, left := 0, 0
	for _, v := range moves {
		if v == 'U' {
			up++
		} else if v == 'D' {
			up--
		} else if v == 'L' {
			left++
		} else if v == 'R' {
			left--
		}
	}

	return up == 0 && left == 0
}

func test01() {
	moves := "UD"
	succResult := true
	fmt.Println("test01 moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", judgeCircle(moves))
	fmt.Println()
}

func test02() {
	moves := "LL"
	succResult := false
	fmt.Println("test02 moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", judgeCircle(moves))
	fmt.Println()
}

func test03() {
	moves := "LLU"
	succResult := false
	fmt.Println("test03 moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", judgeCircle(moves))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
