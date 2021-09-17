package main

import (
	"fmt"
)

func backspaceCompare(s string, t string) bool {
	fmt.Println(backspace(s), backspace(t))
	return backspace(s) == backspace(t)
}

func backspace(str string) string {
	s := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != '#' {
			s = append([]byte{str[i]}, s...)
			continue
		}

		bs_num, s_num := 0, 0
		for j := i; j >= 0; j-- {
			if bs_num > 0 && bs_num <= s_num {
				break
			}
			if str[j] == '#' {
				bs_num++
			} else {
				s_num++
			}
		}
		i -= bs_num*2 - 1
	}

	return string(s)
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

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
}
