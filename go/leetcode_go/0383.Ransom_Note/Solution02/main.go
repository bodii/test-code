package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteLen := len(ransomNote)
	magazineLen := len(magazine)

	if ransomNoteLen > magazineLen {
		return false
	}

	note1 := make(map[byte]int)

	for i := 0; i < magazineLen; i++ {
		note1[magazine[i]] += 1
	}

	for i := 0; i < ransomNoteLen; i++ {
		if v, ok := note1[ransomNote[i]]; !ok || v < 1 {
			return false
		}

		note1[ransomNote[i]] -= 1
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
