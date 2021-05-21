package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	sArr := strings.Split(s, " ")

	words := make([]string, len(sArr))
	for _, v := range sArr {
		strLen := len(v)
		str, indStr := v[:strLen-1], v[strLen-1:strLen]
		ind, _ := strconv.Atoi(string(indStr))
		words[ind-1] = str
	}

	return strings.Join(words, " ")
}

func test01() {
	s := "is2 sentence4 This1 a3"
	succResult := "This is a sentence"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortSentence(s))
	fmt.Println()
}

func test02() {
	s := "Myself2 Me1 I4 and3"
	succResult := "Me Myself and I"
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortSentence(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
