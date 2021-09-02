package main

import (
	"fmt"
)

func reformatNumber(number string) string {
	newNumber := make([]byte, 0)
	for i := 0; i < len(number); i++ {
		if number[i] >= '0' && number[i] <= '9' {
			newNumber = append(newNumber, number[i])
		}
	}

	l := len(newNumber) - 1
	result := make([]byte, 0)
	for i := 0; i <= l; i += 3 {
		if l-i == 3 {
			result = append(result, newNumber[i], newNumber[i+1], '-', newNumber[i+2], newNumber[i+3])
			break
		}

		if i+3 <= l {
			result = append(result, newNumber[i:i+3]...)
		} else {
			result = append(result, newNumber[i:]...)
		}

		if i <= l-3 {
			result = append(result, '-')
		}
	}
	return string(result)
}

func test01() {
	number := "1-23-45 6"
	succResult := "123-456"
	fmt.Println("test01 number:=", number)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reformatNumber(number))
	fmt.Println()
}

func test02() {
	number := "123 4-567"
	succResult := "123-45-67"
	fmt.Println("test02 number:=", number)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reformatNumber(number))
	fmt.Println()
}

func test03() {
	number := "123 4-5678"
	succResult := "123-456-78"
	fmt.Println("test03 number:=", number)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reformatNumber(number))
	fmt.Println()
}

func test04() {
	number := "12"
	succResult := "12"
	fmt.Println("test04 number:=", number)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reformatNumber(number))
	fmt.Println()
}

func test05() {
	number := "--17-5 229 35-39475 "
	succResult := "175-229-353-94-75"
	fmt.Println("test05 number:=", number)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reformatNumber(number))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
