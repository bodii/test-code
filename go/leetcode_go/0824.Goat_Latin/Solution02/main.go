package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	result := strings.Builder{}
	addStr := strings.Builder{}
	addStr.WriteString("maa")

	vowels := map[byte]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	sentenceArr := strings.Split(sentence, " ")

	for k, v := range sentenceArr {
		if k > 0 {
			result.WriteString(" ")
		}
		if vowels[v[0]] {
			result.WriteString(v)
			result.WriteString(addStr.String())
		} else {
			result.WriteString(v[1:])
			result.WriteByte(v[0])
			result.WriteString(addStr.String())
		}

		addStr.WriteString("a")
	}

	return result.String()
}

func test01() {
	s := "I speak Goat Latin"
	succResult := "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", toGoatLatin(s))
	fmt.Println()
}

func test02() {
	s := "The quick brown fox jumped over the lazy dog"
	succResult := "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", toGoatLatin(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
