package main

import "fmt"

func reverseOnlyLetters(S string) string {
	sLen := len(S)

	rev := []byte(S)
	i, j := 0, sLen-1
	for j > i {
		if isEnglishChar(S[i]) && isEnglishChar(S[j]) {
			rev[i] = S[j]
			rev[j] = S[i]
			i++
			j--
			continue
		}

		if !isEnglishChar(S[i]) {
			i++
		}

		if !isEnglishChar(S[j]) {
			j--
		}
	}

	return string(rev)
}

func isEnglishChar(e byte) bool {
	if (97 <= e && e <= 122) || (65 <= e && e <= 90) {
		return true
	}

	return false
}

func test01() {
	S := "ab-cd"
	success := "dc-ba"

	fmt.Println("test01 S:=", S)
	fmt.Println("success:=", success)
	fmt.Println("result:=", reverseOnlyLetters(S))
	fmt.Println()
}

func test02() {
	S := "a-bC-dEf-ghIj"
	success := "j-Ih-gfE-dCba"

	fmt.Println("test02 S:=", S)
	fmt.Println("success:=", success)
	fmt.Println("result:=", reverseOnlyLetters(S))
	fmt.Println()
}

func test03() {
	S := "Test1ng-Leet=code-Q!"
	success := "Qedo1ct-eeLg=ntse-T!"

	fmt.Println("test03 S:=", S)
	fmt.Println("success:=", success)
	fmt.Println("result:=", reverseOnlyLetters(S))
	fmt.Println()
}

func test04() {
	S := "$1:V<&NyJ"
	success := "$1:J<&yNV"

	fmt.Println("test04 S:=", S)
	fmt.Println("success:=", success)
	fmt.Println("result:=", reverseOnlyLetters(S))
	fmt.Println()
}

func test05() {
	S := ""
	success := ""

	fmt.Println("test05 S:=", S)
	fmt.Println("success:=", success)
	fmt.Println("result:=", reverseOnlyLetters(S))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
