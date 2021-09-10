package main

import (
	"fmt"
)

func largeGroupPositions(s string) [][]int {
	lastIndex := len(s) - 1

	result := make([][]int, 0)
	for i := 0; i < lastIndex-1; i++ {
		j := i + 1
		for ; j <= lastIndex; j++ {
			if s[i] != s[j] {
				break
			}
		}

		j--
		if j-i+1 >= 3 {
			result = append(result, []int{i, j})
		}

		i = j
	}

	return result
}

func test01() {
	s := "abbxxxxzzy"
	success := [][]int{{3, 6}}

	fmt.Println("test01 s:", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", largeGroupPositions(s))
	fmt.Println()
}

func test02() {
	s := "abc"
	success := [][]int{}

	fmt.Println("test02 s:", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", largeGroupPositions(s))
	fmt.Println()
}

func test03() {
	s := "abcdddeeeeaabbbcd"
	success := [][]int{{3, 5}, {6, 9}, {12, 14}}

	fmt.Println("test03 s:", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", largeGroupPositions(s))
	fmt.Println()
}

func test04() {
	s := "aba"
	success := [][]int{}

	fmt.Println("test04 s:", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", largeGroupPositions(s))
	fmt.Println()
}

func test05() {
	s := "aaa"
	success := [][]int{{0, 2}}

	fmt.Println("test05 s:", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", largeGroupPositions(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
