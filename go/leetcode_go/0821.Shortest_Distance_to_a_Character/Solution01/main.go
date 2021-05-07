package main

import (
	"fmt"
)

func shortestToChar(s string, c byte) []int {
	rest := make([]int, len(s))
	shortestIndexs := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			shortestIndexs = append(shortestIndexs, i)
		}
	}

	index := 0
	for i := 0; i < len(s); i++ {
		fmt.Println(index, shortestIndexs[index])
		x := abs(shortestIndexs[index] - i)
		if index-1 >= 0 {
			rest[i] = min(x, abs(shortestIndexs[index-1]-i))
		} else {
			rest[i] = x
		}

		if s[i] == c && index < len(shortestIndexs)-1 {
			index++
		}
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
