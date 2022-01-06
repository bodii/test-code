package main

import (
	"fmt"
)

func tictactoe(moves [][]int) string {
	if len(moves) < 5 {
		return "Pending"
	}

	winBinarys := []int{7, 56, 73, 84, 146, 273, 292, 448}
	a, b := 0, 0
	for i := range moves {
		if i&1 == 0 {
			a ^= 1 << (moves[i][0]*3 + moves[i][1])
		} else {
			b ^= 1 << (moves[i][0]*3 + moves[i][1])
		}

	}

	for _, w := range winBinarys {
		if a&w == w {
			return "A"
		}

		if b&w == w {
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

func test01() {
	moves := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}
	succResult := "A"
	fmt.Println("test01:")
	fmt.Println("moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", tictactoe(moves))
	fmt.Println()
}

func test02() {
	moves := [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}
	succResult := "B"
	fmt.Println("test02:")
	fmt.Println("moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", tictactoe(moves))
	fmt.Println()
}

func test03() {
	moves := [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}
	succResult := "Draw"
	fmt.Println("test03:")
	fmt.Println("moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", tictactoe(moves))
	fmt.Println()
}

func test04() {
	moves := [][]int{{0, 0}, {1, 1}}
	succResult := "Pending"
	fmt.Println("test04:")
	fmt.Println("moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", tictactoe(moves))
	fmt.Println()
}

func test05() {
	moves := [][]int{{1, 2}, {2, 1}, {1, 0}, {0, 0}, {0, 1}, {2, 0}, {1, 1}}
	succResult := "A"
	fmt.Println("test05:")
	fmt.Println("moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", tictactoe(moves))
	fmt.Println()
}

func test06() {
	moves := [][]int{{2, 0}, {1, 1}, {0, 2}, {2, 1}, {1, 2}, {1, 0}, {0, 0}, {0, 1}}
	succResult := "B"
	fmt.Println("test06:")
	fmt.Println("moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", tictactoe(moves))
	fmt.Println()
}

func test07() {
	moves := [][]int{{0, 0}, {2, 2}, {1, 0}, {2, 0}, {0, 1}, {1, 2}, {1, 1}, {0, 2}}
	succResult := "B"
	fmt.Println("test07:")
	fmt.Println("moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", tictactoe(moves))
	fmt.Println()
}

func test08() {
	moves := [][]int{{0, 0}, {1, 1}, {2, 2}, {1, 0}, {1, 2}, {0, 2}, {2, 1}, {0, 1}, {2, 0}}
	succResult := "A"
	fmt.Println("test08:")
	fmt.Println("moves:=", moves)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", tictactoe(moves))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
}
