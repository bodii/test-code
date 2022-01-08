package main

import (
	"fmt"
	"strings"
)

func removeVowels(s string) string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	sBytes := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if !vowels[s[i]] {
			sBytes.WriteByte(s[i])
		}
	}

	return sBytes.String()
}

func test01() {
	S := "leetcodeisacommunityforcoders"
	success := "ltcdscmmntyfrcdrs"

	fmt.Println("test01 S:=", S)
	fmt.Println("success:=", success)
	fmt.Println("result:=", removeVowels(S))
	fmt.Println()
}

func test02() {
	S := "aeiou"
	success := ""

	fmt.Println("test02 S:=", S)
	fmt.Println("success:=", success)
	fmt.Println("result:=", removeVowels(S))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
