package main

import (
	"fmt"
)

func findJudge(n int, trust [][]int) int {
	votes := make(map[int]int)

	for _, v := range trust {
		votes[v[0]]--
		votes[v[1]]++
	}

	for i := 1; i <= n; i++ {
		if votes[i] >= n-1 {
			return i
		}
	}

	return -1
}

func test01() {
	n := 2
	trust := [][]int{{1, 2}}
	succResult := 2
	fmt.Println("test01 n:=", n, " trust:=", trust)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findJudge(n, trust))
	fmt.Println()
}

func test02() {
	n := 3
	trust := [][]int{{1, 3}, {2, 3}}
	succResult := 3
	fmt.Println("test02 n:=", n, " trust:=", trust)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findJudge(n, trust))
	fmt.Println()
}

func test03() {
	n := 3
	trust := [][]int{{1, 3}, {2, 3}, {3, 1}}
	succResult := -1
	fmt.Println("test03 n:=", n, " trust:=", trust)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findJudge(n, trust))
	fmt.Println()
}

func test04() {
	n := 3
	trust := [][]int{{1, 2}, {2, 3}}
	succResult := -1
	fmt.Println("test04 n:=", n, " trust:=", trust)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findJudge(n, trust))
	fmt.Println()
}

func test05() {
	n := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	succResult := 3
	fmt.Println("test05 n:=", n, " trust:=", trust)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findJudge(n, trust))
	fmt.Println()
}

func test06() {
	n := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}}
	succResult := -1
	fmt.Println("test06 n:=", n, " trust:=", trust)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findJudge(n, trust))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
