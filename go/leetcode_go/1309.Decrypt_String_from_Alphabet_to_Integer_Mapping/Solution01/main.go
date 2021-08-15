package main

import (
	"fmt"
	"strconv"
	"strings"
)

func freqAlphabets(s string) string {
	sLen := len(s)
	result := strings.Builder{}

	i := 0
	for ; i < sLen-2; i++ {
		if s[i+2] == '#' {
			if n, err := strconv.Atoi(s[i : i+2]); err == nil {
				result.WriteByte(byte(n - 1 + 'a'))
				i += 2
			}
		} else {
			result.WriteByte(s[i] - '1' + 'a')
		}
	}

	for ; i < sLen; i++ {
		result.WriteByte(s[i] - '1' + 'a')
	}

	return result.String()
}

func test01() {
	s := "10#11#12"
	succResult := "jkab"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", freqAlphabets(s))
	fmt.Println()
}

func test02() {
	s := "1326#"
	succResult := "acz"
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", freqAlphabets(s))
	fmt.Println()
}

func test03() {
	s := "25#"
	succResult := "y"
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", freqAlphabets(s))
	fmt.Println()
}

func test04() {
	s := "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
	succResult := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", freqAlphabets(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
