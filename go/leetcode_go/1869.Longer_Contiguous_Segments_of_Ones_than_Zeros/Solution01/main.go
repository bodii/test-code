package main

import (
	"fmt"
)

func checkZeroOnes(s string) bool {
	onesMax, zeroMax := 0, 0
	sLen := len(s)

	for k := 0; k < sLen; k++ {
		if s[k] == '1' {
			i := k
			for i < sLen && s[i] == '1' {
				i++
			}

			if i-k > onesMax {
				onesMax = i - k
			}
			k = i - 1
		} else {
			i := k
			for i < sLen && s[i] == '0' {
				i++
			}

			if i-k > zeroMax {
				zeroMax = i - k
			}
			k = i - 1
		}
	}

	return onesMax > zeroMax
}

func test01() {
	s := "1101"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkZeroOnes(s))
	fmt.Println()
}

func test02() {
	s := "111000"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkZeroOnes(s))
	fmt.Println()
}

func test03() {
	s := "110100010"
	succResult := false
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkZeroOnes(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
