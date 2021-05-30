package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sArr := strings.Split(s, " ")

	resltus := make([]string, len(sArr))
	for k, v := range sArr {
		vLen := len(v)
		vBytes := make([]byte, vLen)
		j := 0
		for i := vLen - 1; i >= 0; i-- {
			vBytes[j] = v[i]
			j++
		}
		resltus[k] = string(vBytes)
	}

	return strings.Join(resltus, " ")
}

func test01() {
	s := "Let's take LeetCode contest"
	successResult := "s'teL ekat edoCteeL tsetnoc"
	fmt.Println("test01 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", reverseWords(s))
	fmt.Println()
}

func test02() {
	s := "God Ding"
	successResult := "doG gniD"
	fmt.Println("test02 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", reverseWords(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
