package main

import (
	"fmt"
)

func numberOfLines(widths []int, s string) []int {
	lines, lineSum := 1, 0
	for i := 0; i < len(s); i++ {
		num := widths[int(s[i]-'a')]
		if lineSum+num > 100 {
			lines++
			lineSum = num
		} else {
			lineSum += num
		}
	}

	return []int{lines, lineSum}
}

func test01() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"
	succResult := []int{3, 60}
	fmt.Printf("test01 widths:= %v, s:=%v", widths, s)
	fmt.Printf("success result:  %v \n", succResult)
	fmt.Println("result:", numberOfLines(widths, s))
	fmt.Println()
}

func test02() {
	widths := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "bbbcccdddaaa"
	succResult := []int{2, 4}
	fmt.Printf("test02 widths:= %v, s:=%v", widths, s)
	fmt.Printf("success result:  %v \n", succResult)
	fmt.Println("result:", numberOfLines(widths, s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
