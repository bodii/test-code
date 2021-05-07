package main

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) []int {
	rest := make([]int, len(s))
	len := len(s)

	prev := math.MaxInt32 / 2
	for i := 0; i < len; i++ {
		if s[i] == c {
			prev = i
		}
		rest[i] = i - prev
	}

	prev = math.MaxInt32 / 2
	for i := len - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		fmt.Println(rest[i], prev-i)
		rest[i] = min(abs(rest[i]), abs(prev-i))
	}

	return rest
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func test01() {
	s := "loveleetcode"
	c := "e"
	reslut := []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}
	fmt.Printf("test01 s:= %s, c:= %s\n", s, c)
	fmt.Printf("success result:  %v \n", reslut)
	fmt.Println("result:", shortestToChar(s, []byte(c)[0]))
	fmt.Println()
}

func test02() {
	s := "aaab"
	c := "b"
	reslut := []int{3, 2, 1, 0}
	fmt.Printf("test02 s:= %s, c:= %s\n", s, c)
	fmt.Printf("success result:  %v \n", reslut)
	fmt.Println("result:", shortestToChar(s, []byte(c)[0]))
	fmt.Println()
}

func test03() {
	s := "aaba"
	c := "b"
	reslut := []int{2, 1, 0, 1}
	fmt.Printf("test03 s:= %s, c:= %s\n", s, c)
	fmt.Printf("success result:  %v \n", reslut)
	fmt.Println("result:", shortestToChar(s, []byte(c)[0]))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
