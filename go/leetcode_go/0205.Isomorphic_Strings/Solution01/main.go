package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	sMap, tMap := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		t1, t1ok := sMap[s[i]]
		s1, s1ok := tMap[t[i]]
		if !t1ok && !s1ok {
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		} else if (!t1ok || !s1ok) || (t1ok && s1ok && (t1 != t[i] || s1 != s[i])) {
			return false
		}
	}

	return true
}

func test01() {
	s, t := "egg", "add"
	succResult := true
	fmt.Println("test01 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isIsomorphic(s, t))
	fmt.Println()
}

func test02() {
	s, t := "foo", "bar"
	succResult := false
	fmt.Println("test02 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isIsomorphic(s, t))
	fmt.Println()
}

func test03() {
	s, t := "paper", "title"
	succResult := true
	fmt.Println("test03 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isIsomorphic(s, t))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
