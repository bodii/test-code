package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteLen := len(ransomNote)
	magazineLen := len(magazine)

	if ransomNoteLen > magazineLen {
		return false
	}

	note1 := make(map[byte]int)
	note2 := make(map[byte]int)

	for i := 0; i < magazineLen; i++ {
		note1[magazine[i]] += 1
	}

	for i := 0; i < ransomNoteLen; i++ {
		note2[ransomNote[i]] += 1
	}

	for k, v2 := range note2 {
		v1, ok := note1[k]
		if !ok {
			return false
		}

		if v2 > v1 {
			return false
		}
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

func main() {
	test01()
	test02()
	test03()
	test04()
}
