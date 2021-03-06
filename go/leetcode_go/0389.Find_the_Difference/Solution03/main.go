package main

import "fmt"

func findTheDifference(s string, t string) byte {
	sum := 0
	sLen := len(s)

	for i := 0; i <= sLen; i++ {
		sum += int(t[i])
		if i < sLen {
			sum -= int(s[i])
		}
	}

	return byte(sum)
}

func test01() {
	s := "abcd"
	t := "abcde"
	fmt.Println("test01 result:", findTheDifference(s, t))
}

func test02() {
	s := ""
	t := "y"
	fmt.Println("test02 result:", findTheDifference(s, t))
}

func test03() {
	s := "a"
	t := "aa"
	fmt.Println("test03 result:", findTheDifference(s, t))
}

func test04() {
	s := "ae"
	t := "aea"
	fmt.Println("test04 result:", findTheDifference(s, t))
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
