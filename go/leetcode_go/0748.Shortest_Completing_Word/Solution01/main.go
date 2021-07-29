package main

import (
	"fmt"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlateMap := toMap(licensePlate)

	minLen, curWord := 15, ""
	for _, word := range words {
		wordMap := toMap(word)

		ok, len := contain(licensePlateMap, wordMap)
		if !ok {
			continue
		}

		if minLen > len {
			minLen = len
			curWord = word
		}
	}
	return curWord
}

func contain(src, a map[byte]int) (isOk bool, len int) {
	isOk, len = false, 0
	for k, v := range src {
		if v1, ok := a[k]; !ok || v1 < v {
			return
		}
	}

	isOk = true
	for _, v := range a {
		len += v
	}

	return
}

func toMap(str string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			s := str[i]
			if s >= 'A' && s <= 'Z' {
				s += 'a' - 'A'
			}
			m[s]++
		}
	}

	return m
}

func test01() {
	licensePlate := "1s3 PSt"
	words := []string{"step", "steps", "stripe", "stepple"}
	succResult := "steps"
	fmt.Println("test01 licensePlate:=", licensePlate, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", shortestCompletingWord(licensePlate, words))
	fmt.Println()
}

func test02() {
	licensePlate := "1s3 456"
	words := []string{"looks", "pest", "stew", "show"}
	succResult := "pest"
	fmt.Println("test02 licensePlate:=", licensePlate, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", shortestCompletingWord(licensePlate, words))
	fmt.Println()
}

func test03() {
	licensePlate := "Ah71752"
	words := []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}
	succResult := "husband"
	fmt.Println("test03 licensePlate:=", licensePlate, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", shortestCompletingWord(licensePlate, words))
	fmt.Println()
}

func test04() {
	licensePlate := "OgEu755"
	words := []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}
	succResult := "enough"
	fmt.Println("test04 licensePlate:=", licensePlate, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", shortestCompletingWord(licensePlate, words))
	fmt.Println()
}

func test05() {
	licensePlate := "iMSlpe4"
	words := []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"}
	succResult := "simple"
	fmt.Println("test05 licensePlate:=", licensePlate, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", shortestCompletingWord(licensePlate, words))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
