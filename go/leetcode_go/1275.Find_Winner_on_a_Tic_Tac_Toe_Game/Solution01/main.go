package main

import (
	"fmt"
	"sort"
	"strconv"
)

func tictactoe(moves [][]int) string {
	if len(moves) < 5 {
		return "Pending"
	}

	wins := [][]string{{"00", "11", "22"}, {"00", "01", "02"},
		{"00", "10", "20"}, {"01", "11", "21"}, {"10", "11", "12"},
		{"02", "11", "20"}, {"02", "12", "22"}, {"20", "21", "22"}}

	A, B := make([]string, 0), make([]string, 0)
	for k, v := range moves {
		if k&1 == 0 {
			A = append(A, strconv.Itoa(v[0])+strconv.Itoa(v[1]))
		} else {
			B = append(B, strconv.Itoa(v[0])+strconv.Itoa(v[1]))
		}
	}

	// fmt.Println(A)
	aIsWin, bIsWin := isWin(wins, A), isWin(wins, B)

	if aIsWin == bIsWin && (aIsWin || (!aIsWin && len(moves) == 9)) {
		return "Draw"
	} else if aIsWin && !bIsWin {
		return "A"
	} else if !aIsWin && bIsWin {
		return "B"
	}

	return "Pending"
}

func isWin(wins [][]string, member []string) bool {
	if len(member) < 3 {
		return false
	}

	sort.SliceStable(member, func(i, j int) bool {
		return member[i] < member[j]
	})
	fmt.Println(member)

	winNum := 0
	for _, win := range wins {
		if winNum == 3 {
			return true
		}

		winNum = 0
		for _, v := range member {
			if winNum == 3 {
				return true
			}
			if win[winNum] == v {
				winNum++
			}
		}
	}
	fmt.Println(winNum, wins)

	return winNum == 3
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
