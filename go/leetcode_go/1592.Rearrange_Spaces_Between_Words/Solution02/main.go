package main

import (
	"fmt"
	"strings"
)

func reorderSpaces(text string) string {
	spaces := strings.Count(text, " ")
	words := strings.Fields(text)

	if spaces > 0 {
		if len(words) > 1 {
			spaceStr := strings.Repeat(" ", spaces/(len(words)-1))
			text = strings.Join(words, spaceStr)
			text += strings.Repeat(" ", spaces%(len(words)-1))
		} else {
			text = words[0] + strings.Repeat(" ", spaces)
		}
	}

	return text
}

func test01() {
	s := "  this   is  a sentence "
	succResult := "[this   is   a   sentence]"

	fmt.Println("test01:")
	fmt.Println("s:=", s)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", "["+reorderSpaces(s)+"]")
	fmt.Println()
}

func test02() {
	s := " practice   makes   perfect"
	succResult := "[practice   makes   perfect ]"

	fmt.Println("test02:")
	fmt.Println("s:=", s)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", "["+reorderSpaces(s)+"]")
	fmt.Println()
}

func test03() {
	s := "hello   world"
	succResult := "[hello   world]"

	fmt.Println("test03:")
	fmt.Println("s:=", s)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", "["+reorderSpaces(s)+"]")
	fmt.Println()
}

func test04() {
	s := "  walks  udp package   into  bar a"
	succResult := "[walks  udp  package  into  bar  a ]"

	fmt.Println("test04:")
	fmt.Println("s:=", s)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", "["+reorderSpaces(s)+"]")
	fmt.Println()
}

func test05() {
	s := "a"
	succResult := "[a]"

	fmt.Println("test05:")
	fmt.Println("s:=", s)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", "["+reorderSpaces(s)+"]")
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
