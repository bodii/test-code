package main

import (
	"fmt"
)

func makeGood(s string) string {
	sLen := len(s)
	resultStack := newStack(sLen)

	for i := 0; i < sLen; i++ {
		if !resultStack.isEmpty() && toUpper(resultStack.get()) == toUpper(s[i]) && resultStack.get() != s[i] {
			resultStack.pop()
		} else {
			resultStack.push(s[i])
		}
	}

	if resultStack.isEmpty() {
		return ""
	}

	return string(resultStack.getAll())
}

type stack struct {
	top  int
	cap  int
	data []byte
}

func (s *stack) push(a byte) {
	s.top++
	s.data[s.top] = a
}

func (s *stack) pop() byte {
	sVal := s.data[s.top]
	s.top--
	return sVal
}

func (s stack) get() byte {
	return s.data[s.top]
}

func (s stack) getSize() int {
	return s.top + 1
}

func (s stack) isEmpty() bool {
	return s.getSize() == 0
}

func (s stack) getAll() []byte {
	return s.data[:s.getSize()]
}

func newStack(cap int) *stack {
	return &stack{
		top:  -1,
		cap:  cap,
		data: make([]byte, cap),
	}
}

func toUpper(a byte) byte {
	if isLowerChar(a) {
		return a - 32
	}

	return a
}

func isLowerChar(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}

	return false
}

func test01() {
	s := "leEeetcode"
	succResult := "leetcode"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeGood(s))
	fmt.Println()
}

func test02() {
	s := "abBAcC"
	succResult := ""
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeGood(s))
	fmt.Println()
}

func test03() {
	s := "s"
	succResult := "s"
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeGood(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
