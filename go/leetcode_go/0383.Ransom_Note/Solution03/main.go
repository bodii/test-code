package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	note := make(map[rune]int, len(magazine))
	for _, v := range magazine {
		note[v] += 1
	}
	for _, v := range ransomNote {
		if c, ok := note[v]; !ok || c < 1 {
			return false
		}
		note[v] -= 1
	}

	return true
}

func test01() {
	ransomNote, magazine := "a", "b"

	fmt.Println("test 01 the result: ", canConstruct(ransomNote, magazine))
}

func test02() {
	ransomNote, magazine := "aa", "ab"

	fmt.Println("test 02 the result: ", canConstruct(ransomNote, magazine))
}

func test03() {
	ransomNote, magazine := "aa", "aab"

	fmt.Println("test 03 the result: ", canConstruct(ransomNote, magazine))
}

func test04() {
	ransomNote, magazine := "bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"

	fmt.Println("test 04 the result: ", canConstruct(ransomNote, magazine))
}

func test05() {
	ransomNote, magazine := "aab", "baa"

	fmt.Println("test 05 the result: ", canConstruct(ransomNote, magazine))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
