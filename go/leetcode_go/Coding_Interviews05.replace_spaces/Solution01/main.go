package main

import (
	"fmt"
	"regexp"
)

func replaceSpace(s string) string {
	return regexp.MustCompile(`\s`).ReplaceAllString(s, "%20")
}

func test01() {
	s := "We are happy."
	succResult := "We%20are%20happy."
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceSpace(s))
	fmt.Println()
}

func main() {
	test01()
}
