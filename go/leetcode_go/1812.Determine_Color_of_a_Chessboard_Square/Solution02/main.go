package main

import (
	"fmt"
)

func squareIsWhite(coordinates string) bool {
	return (coordinates[0]-'a'+coordinates[1])&1 == 0
}

func test01() {
	s := "a1"
	successResult := false
	fmt.Println("test01  s: ", s)
	fmt.Println("success result: ", successResult)
	fmt.Println("result: ", squareIsWhite(s))
	fmt.Println()
}

func test02() {
	s := "h3"
	successResult := true
	fmt.Println("test01  s: ", s)
	fmt.Println("success result: ", successResult)
	fmt.Println("result: ", squareIsWhite(s))
	fmt.Println()
}

func test03() {
	s := "c7"
	successResult := false
	fmt.Println("test01  s: ", s)
	fmt.Println("success result: ", successResult)
	fmt.Println("result: ", squareIsWhite(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
