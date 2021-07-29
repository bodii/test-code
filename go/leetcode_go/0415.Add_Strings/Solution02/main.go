package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	result := make([]byte, 0)
	num1Len, num2Len := len(num1), len(num2)

	rem := 0
	j1, j2 := num1Len, num2Len
	for j1 > 0 || j2 > 0 {
		if j1 > 0 {
			j1--
			rem += int(num1[j1] - '0')
		}

		if j2 > 0 {
			j2--
			rem += int(num2[j2] - '0')
		}

		cur := rem % 10
		if rem >= 10 {
			rem /= 10
		} else {
			rem = 0
		}

		result = append([]byte{byte(cur + '0')}, result...)
	}

	if rem > 0 {
		result = append([]byte{byte(rem + '0')}, result...)
	}

	return string(result)
}

func test01() {
	num1 := "11"
	num2 := "123"
	succResult := "134"
	fmt.Println("test01 num1:=", num1, " num2:=", num2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addStrings(num1, num2))
	fmt.Println()
}

func test02() {
	num1 := "456"
	num2 := "77"
	succResult := "533"
	fmt.Println("test02 num1:=", num1, " num2:=", num2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addStrings(num1, num2))
	fmt.Println()
}

func test03() {
	num1 := "0"
	num2 := "0"
	succResult := "0"
	fmt.Println("test03 num1:=", num1, " num2:=", num2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addStrings(num1, num2))
	fmt.Println()
}

func test04() {
	num1 := "3876620623801494171"
	num2 := "6529364523802684779"
	succResult := "10405985147604178950"
	fmt.Println("test04 num1:=", num1, " num2:=", num2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addStrings(num1, num2))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
