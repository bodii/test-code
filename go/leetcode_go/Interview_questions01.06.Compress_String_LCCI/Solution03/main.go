package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	SLen := len(S)
	if SLen == 0 {
		return S
	}

	result := []byte{}
	cur, curNum := S[0], 1
	for i := 1; i < len(S); i++ {
		if S[i] == cur {
			curNum++
		} else {
			result = append(result, cur)
			result = append(result, []byte(strconv.Itoa(curNum))...)
			cur = S[i]
			curNum = 1
		}
	}

	result = append(result, cur)
	result = append(result, []byte(strconv.Itoa(curNum))...)

	if len(S) <= len(result) {
		return S
	}

	return string(result)
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
