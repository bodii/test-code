package main

import (
	"fmt"
)

func maxDepth(s string) int {
	maxDepth := 0
	stackDepth := newStack(len(s) / 2)

	for _, v := range s {
		if v == '(' {
			stackDepth.push(1)
			if stackDepth.getSize() > maxDepth {
				maxDepth = stackDepth.getSize()
			}
		} else if v == ')' {
			stackDepth.pop()
		}
	}

	return maxDepth
}

type stack struct {
	top  int
	cap  int
	data []int
}

func (s *stack) push(a int) {
	s.top++
	s.data[s.top] = a
}

func (s *stack) pop() int {
	sVal := s.data[s.top]
	s.top--
	return sVal
}

func (s stack) peek() int {
	return s.data[s.top]
}

func (s stack) getSize() int {
	return s.top + 1
}

func (s stack) isEmpty() bool {
	return s.getSize() == 0
}

func (s stack) getAll() []int {
	return s.data[:s.getSize()]
}

func newStack(cap int) *stack {
	return &stack{
		top:  -1,
		cap:  cap,
		data: make([]int, cap),
	}
}

func test01() {
	s := "(1+(2*3)+((8)/4))+1"
	succResult := 3
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func test02() {
	s := "(1)+((2))+(((3)))"
	succResult := 3
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func test03() {
	s := "1+(2*3)/(2-1)"
	succResult := 1
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func test04() {
	s := "1"
	succResult := 0
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func test05() {
	s := "8*((1*(5+6))*(8/6))"
	succResult := 3
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
