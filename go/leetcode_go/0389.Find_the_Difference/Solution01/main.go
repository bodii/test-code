package main

import "fmt"

func findTheDifference(s string, t string) byte {
	tLen := len(t)
	sLen := len(s)
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < tLen; i++ {
		if i < sLen {
			sMap[s[i]] += 1
		}
		tMap[t[i]] += 1
	}

	var result byte
	for k, tNum := range tMap {
		sNum, ok := sMap[k]
		if !ok || sNum != tNum {
			result = k
		}
	}

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
