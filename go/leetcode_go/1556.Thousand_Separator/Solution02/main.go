package main

import (
	"fmt"
	"strconv"
)

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	result := []byte{}
	for i := 0; n > 0; i++ {
		if i > 0 && i%3 == 0 {
			result = append([]byte{'.'}, result...)
		}
		result = append([]byte{strconv.Itoa(n % 10)[0]}, result...)
		n /= 10
	}

	return string(result)
}

func test01() {
	n := 987
	succResult := "987"
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", thousandSeparator(n))
	fmt.Println()
}

func test02() {
	n := 1234
	succResult := "1.234"
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", thousandSeparator(n))
	fmt.Println()
}

func test03() {
	n := 123456789
	succResult := "123.456.789"
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", thousandSeparator(n))
	fmt.Println()
}

func test04() {
	n := 0
	succResult := "0"
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", thousandSeparator(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
