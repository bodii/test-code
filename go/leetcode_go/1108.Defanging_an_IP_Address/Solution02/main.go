package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	result := strings.Builder{}
	for _, v := range address {
		if v != '.' {
			result.WriteRune(v)
		} else {
			result.WriteString("[.]")
		}
	}

	return result.String()
}

func test01() {
	address := "1.1.1.1"
	succResult := "1[.]1[.]1[.]1"
	fmt.Println("test01 address:=", address)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", defangIPaddr(address))
	fmt.Println()
}

func test02() {
	address := "255.100.50.0"
	succResult := "255[.]100[.]50[.]0"
	fmt.Println("test02 address:=", address)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", defangIPaddr(address))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
