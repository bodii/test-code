package main

import "fmt"

func findTheDifference(s string, t string) byte {
	tLen := len(t)
	tMap := make(map[byte]int)

	for i := 0; i < tLen; i++ {
		tMap[t[i]] += 1
		if i < tLen-1 {
			tMap[s[i]] -= 1
		}
	}

	for k, v := range tMap {
		if v > 0 {
			return k
		}
	}

	var result byte
	return result
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
