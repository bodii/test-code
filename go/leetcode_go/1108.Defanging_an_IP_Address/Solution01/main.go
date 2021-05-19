package main

import (
	"fmt"
)

func defangIPaddr(address string) string {
	b := []byte(address)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] != '.' {
			result = append(result, b[i])
		} else {
			result = append(result, '[', '.', ']')
		}
	}

	return string(result)
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
