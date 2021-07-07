package main

import (
	"fmt"
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	digitals := make([]string, 0)
	letters := make([]string, 0)
	for _, log := range logs {
		content := strings.Split(log, " ")
		if 97 > content[1][0] {
			digitals = append(digitals, log)
		} else {
			letters = append(letters, log)
		}
	}

	// sort
	sort.Slice(letters, func(i, j int) bool {
		c1 := strings.SplitN(letters[i], " ", 2)
		c2 := strings.SplitN(letters[j], " ", 2)
		if (c1[1] == c2[1] && c1[0] > c2[0]) || c1[1] > c2[1] {
			return false
		}
		return true
	})

	return append(letters, digitals...)
}

func test01() {
	logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	succResult := []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"}
	fmt.Println("test01 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reorderLogFiles(logs))
	fmt.Println()
}

func test02() {
	logs := []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}
	succResult := []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"}
	fmt.Println("test02 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reorderLogFiles(logs))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
