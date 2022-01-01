package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	maxIdn, minIdn := len(a)-1, len(b)-1
	max, min := a, b
	if maxIdn < minIdn {
		maxIdn, minIdn = minIdn, maxIdn
		max, min = min, max
	}

	resultBytes := make([]byte, 0)
	rem := 0
	for i := 0; i <= maxIdn; i++ {
		val := rem
		if max[maxIdn-i] == '1' {
			val += 1
		}

		if minIdn-i >= 0 && min[minIdn-i] == '1' {
			val += 1
		}

		rem, val = val>>1, val&1

		resultBytes = append([]byte{byte('0' + val)}, resultBytes...)
	}

	if rem > 0 {
		resultBytes = append([]byte{byte('0' + 1)}, resultBytes...)
	}

	return string(resultBytes)
}

func test01() {
	a := "11"
	b := "10"
	succResult := "101"
	fmt.Println("test01 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addBinary(a, b))
	fmt.Println()
}

func test02() {
	a := "1010"
	b := "1011"
	succResult := "10101"
	fmt.Println("test02 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addBinary(a, b))
	fmt.Println()
}

func test03() {
	a := "101011111111111111111011111100110000110001110011000"
	b := "1011000110010011000110101000110011100011000100100011001101"
	succResult := "1011001011110011000110101000010011001001001010110001100101"
	fmt.Println("test3 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addBinary(a, b))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
