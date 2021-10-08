package main

import (
	"fmt"
)

func replaceSpaces(S string, length int) string {
	str := []byte(S)
	result := []byte{}
	for k, v := range str {
		if k+1 > length {
			return string(result)
		}

		if v == ' ' {
			result = append(result, '%', '2', '0')
		} else {
			result = append(result, v)
		}
	}

	return string(result)
}

func test01() {
	S, length := "Mr John Smith    ", 13
	succResult := "Mr%20John%20Smith"
	fmt.Println("test01 S:=", S, " length:=", length)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceSpaces(S, length))
	fmt.Println()
}

func test02() {
	S, length := "               ", 5
	succResult := "%20%20%20%20%20"
	fmt.Println("test02 S:=", S, " length:=", length)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceSpaces(S, length))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
