package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	addStr := "maa"
	vowels := map[byte]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	sentenceArr := strings.Split(sentence, " ")

	for k, v := range sentenceArr {
		if vowels[v[0]] {
			sentenceArr[k] = v + addStr
		} else {
			sentenceArr[k] = fmt.Sprintf("%s%s%s", v[1:], string(v[0]), addStr)
		}
		addStr += "a"
	}

	return strings.Join(sentenceArr, " ")
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
