package main

import (
	"fmt"
)

func compressString(S string) string {
	SLen := len(S)
	result := []byte{}
	for i := 0; i < SLen; i++ {
		j := i + 1
		for ; j < SLen; j++ {
			if S[i] != S[j] {
				break
			}
		}

		result = append(result, S[i])
		result = append(result, intToBytes(j-i)...)
		i = j - 1
	}

	if len(S) <= len(result) {
		return S
	}

	return string(result)
}

func intToBytes(n int) []byte {
	res := []byte{}

	for n > 0 {
		res = append([]byte{byte('0' + (n % 10))}, res...)
		n /= 10
	}

	return res
}

func test01() {
	s := "aabcccccaaa"
	successResult := "a2b1c5a3"
	fmt.Println("test01 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("        result:", compressString(s))
	fmt.Println()
}

func test02() {
	s := "abbccd"
	successResult := "abbccd"
	fmt.Println("test02 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("        result:", compressString(s))
	fmt.Println()
}

func test03() {
	s := "bb"
	successResult := "bb"
	fmt.Println("test03 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("        result:", compressString(s))
	fmt.Println()
}

func test04() {
	s := "rrrrrLLLLLPPPPPPRRRRRgggNNNNNVVVVVVVVVVDDDDDDDDDDIIIIIIIIIIlllllllAAAAqqqqqqqbbbNNNNffffff"
	successResult := "r5L5P6R5g3N5V10D10I10l7A4q7b3N4f6"
	fmt.Println("test04 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("        result:", compressString(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
