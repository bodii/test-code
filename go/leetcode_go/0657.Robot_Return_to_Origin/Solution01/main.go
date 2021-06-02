package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	dirs := map[rune]int{'U': 1, 'D': -1, 'L': 1, 'R': -1}
	up, left := 0, 0

	for _, d := range moves {
		if d == 'U' || d == 'D' {
			up += dirs[d]
		} else {
			left += dirs[d]
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

func main() {
	test01()
	test02()
}
