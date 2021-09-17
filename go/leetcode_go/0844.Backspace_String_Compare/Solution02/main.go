package main

import (
	"fmt"
)

func backspaceCompare(s string, t string) bool {
	sLen, tLen := len(s)-1, len(t)-1

	for sLen >= 0 || tLen >= 0 {
		sLen = skip(s, sLen)
		tLen = skip(t, tLen)

		if ((sLen < 0 || tLen < 0) && sLen != tLen) || (sLen >= 0 && tLen >= 0 && s[sLen] != t[tLen]) {
			return false
		}

		sLen--
		tLen--
	}

	return true
}

func skip(s string, index int) int {
	skipNum := 0
	for index >= 0 {
		if s[index] == '#' {
			skipNum++
			index--
		} else if skipNum > 0 {
			skipNum--
			index--
		} else {
			return index
		}
	}

	return index
}

func test01() {
	s, t := "ab#c", "ad#c"
	success := true

	fmt.Println("test01 s:=", s, " t:=", t)
	fmt.Println("success result:", success)
	fmt.Println("result:", backspaceCompare(s, t))
	fmt.Println()
}

func test02() {
	s, t := "ab##", "c#d#"
	success := true

	fmt.Println("test02 s:=", s, " t:=", t)
	fmt.Println("success result:", success)
	fmt.Println("result:", backspaceCompare(s, t))
	fmt.Println()
}

func test03() {
	s, t := "a##c", "#a#c"
	success := true

	fmt.Println("test03 s:=", s, " t:=", t)
	fmt.Println("success result:", success)
	fmt.Println("result:", backspaceCompare(s, t))
	fmt.Println()
}

func test04() {
	s, t := "a#c", "b"
	success := false

	fmt.Println("test04 s:=", s, " t:=", t)
	fmt.Println("success result:", success)
	fmt.Println("result:", backspaceCompare(s, t))
	fmt.Println()
}

func test05() {
	s, t := "bxj##tw", "bxo#j##tw"
	success := true

	fmt.Println("test05 s:=", s, " t:=", t)
	fmt.Println("success result:", success)
	fmt.Println("result:", backspaceCompare(s, t))
	fmt.Println()
}

func test06() {
	s, t := "xywrrmp", "xywrrmu#p"
	success := true

	fmt.Println("test06 s:=", s, " t:=", t)
	fmt.Println("success result:", success)
	fmt.Println("result:", backspaceCompare(s, t))
	fmt.Println()
}

func test07() {
	s, t := "nzp#o#g", "b#nzp#o#g"
	success := true

	fmt.Println("test07 s:=", s, " t:=", t)
	fmt.Println("success result:", success)
	fmt.Println("result:", backspaceCompare(s, t))
	fmt.Println()
}

func test08() {
	s, t := "bxj##tw", "bxj###tw"
	success := false

	fmt.Println("test08 s:=", s, " t:=", t)
	fmt.Println("success result:", success)
	fmt.Println("result:", backspaceCompare(s, t))
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
