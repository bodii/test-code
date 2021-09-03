package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	l := len(s)
	result := []byte(s)

	for i := 0; i < l; i += k * 2 {
		right := i + k
		if right > l {
			right = l
		}
		reverse(&result, i, right-1)
	}

	return string(result)
}

func reverse(s *[]byte, left, right int) {
	for left < right {
		(*s)[left], (*s)[right] = (*s)[right], (*s)[left]
		left++
		right--
	}
}

func test01() {
	s := "abcdefg"
	k := 2
	successResult := "bacdfeg"
	fmt.Println("test01 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", reverseStr(s, k))
	fmt.Println()
}

func test02() {
	s := "abcd"
	k := 2
	successResult := "bacd"
	fmt.Println("test02 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", reverseStr(s, k))
	fmt.Println()
}

func test03() {
	s := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
	k := 39
	successResult := "fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi"
	fmt.Println("test03 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", reverseStr(s, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
