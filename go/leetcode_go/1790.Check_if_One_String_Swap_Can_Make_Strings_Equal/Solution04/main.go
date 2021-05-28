package main

import (
	"fmt"
)

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	myStack := newStack(2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if myStack.size() == 1 {
				return false
			} else if myStack.size() < 1 {
				myStack.push(s1[i])
				myStack.push(s2[i])
			} else if myStack.size() > 1 {
				if myStack.poll() != s2[i] || myStack.peek() != s1[i] {
					return false
				}
			}
		}
	}

	return myStack.size() <= 1
}

type stack struct {
	data []byte
	top  int
}

func newStack(cap int) *stack {
	return &stack{
		data: make([]byte, cap),
		top:  -1,
	}
}

func (s *stack) push(b byte) {
	s.top++
	s.data[s.top] = b
}

func (s *stack) peek() byte {
	if s.size() == 0 {
		return 0
	}

	return s.data[0]
}

func (s *stack) size() int {
	return s.top + 1
}

func (s *stack) empty() bool {
	return s.top >= 0
}

func (s *stack) poll() byte {
	if s.size() == 0 {
		return 0
	}

	v := s.data[0]
	s.data = s.data[1:]
	s.top--
	return v
}

func test01() {
	s1, s2 := "bank", "kanb"
	successResult := true
	fmt.Println("test01 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test02() {
	s1, s2 := "attack", "defend"
	successResult := false
	fmt.Println("test02 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test03() {
	s1, s2 := "kelb", "kelb"
	successResult := true
	fmt.Println("test03 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test04() {
	s1, s2 := "abcd", "dcba"
	successResult := false
	fmt.Println("test04 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test05() {
	s1, s2 := "ysmpagrkzsmmzmsssutzgpxrmoylkgemgfcperptsxjcsgojwourhxlhqkxumonfgrczmjvbhwvhpnocz",
		"ysmpagrqzsmmzmsssutzgpxrmoylkgemgfcperptsxjcsgojwourhxlhkkxumonfgrczmjvbhwvhpnocz"

	successResult := true
	fmt.Println("test05:")
	fmt.Println("s1:=", s1)
	fmt.Println("s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test06() {
	s1, s2 := "ab", "ac"

	successResult := false
	fmt.Println("test06:")
	fmt.Println("s1:=", s1)
	fmt.Println("s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test07() {
	s1, s2 := "qgqeg", "gqgeq"

	successResult := false
	fmt.Println("test07:")
	fmt.Println("s1:=", s1)
	fmt.Println("s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
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
