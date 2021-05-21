package main

import (
	"fmt"
)

func restoreString(s string, indices []int) string {
	resultBytes := make([]byte, len(s))

	for k, v := range indices {
		resultBytes[v] = s[k]
	}

	return string(resultBytes)
}

func test01() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	succResult := "leetcode"
	fmt.Println("test01 s:=", s, " indices:=", indices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", restoreString(s, indices))
	fmt.Println()
}

func test02() {
	s := "abc"
	indices := []int{0, 1, 2}
	succResult := "abc"
	fmt.Println("test02 s:=", s, " indices:=", indices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", restoreString(s, indices))
	fmt.Println()
}

func test03() {
	s := "aiohn"
	indices := []int{3, 1, 4, 2, 0}
	succResult := "nihao"
	fmt.Println("test3 s:=", s, " indices:=", indices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", restoreString(s, indices))
	fmt.Println()
}

func test04() {
	s := "aaiougrt"
	indices := []int{4, 0, 2, 6, 7, 3, 1, 5}
	succResult := "arigatou"
	fmt.Println("test04 s:=", s, " indices:=", indices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", restoreString(s, indices))
	fmt.Println()
}

func test05() {
	s := "art"
	indices := []int{1, 0, 2}
	succResult := "rat"
	fmt.Println("test05 s:=", s, " indices:=", indices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", restoreString(s, indices))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
