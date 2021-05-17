package main

import (
	"fmt"
)

func checkIfPangram(sentence string) bool {
	sentenceLen := len(sentence)
	if sentenceLen < 26 {
		return false
	}

	strs := make([]int, 26)
	mid := sentenceLen / 2
	for i := 0; i < mid; i++ {
		strs[int(sentence[i]-'a')]++
		strs[int(sentence[sentenceLen-1-i]-'a')]++
	}

	if sentenceLen%2 != 0 {
		strs[int(sentence[mid+1]-'a')]++
	}

	for _, v := range strs {
		if v == 0 {
			return false
		}
	}

	return true
}

func test01() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	succResult := true
	fmt.Println("test01 sentence:=", sentence)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkIfPangram(sentence))
	fmt.Println()
}

func test02() {
	sentence := "leetcode"
	succResult := false
	fmt.Println("test01 gain:=", sentence)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkIfPangram(sentence))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
