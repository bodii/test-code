package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	up, left := 0, 0
	i, j := 0, len(moves)-1

	for i <= j {
		dir(moves[i], &up, &left)
		i++
		if i < j {
			dir(moves[j], &up, &left)
			j--
		}
	}

	return up == 0 && left == 0
}

func dir(m byte, up, left *int) {
	if m == 'U' {
		(*up)++
	} else if m == 'D' {
		(*up)--
	} else if m == 'L' {
		(*left)++
	} else if m == 'R' {
		(*left)--
	}
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
