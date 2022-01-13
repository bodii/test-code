package main

import (
	"fmt"
)

func capitalizeTitle(title string) string {
	titles := []byte(title)

	titles = append(titles, ' ')
	titlesLastIdn, j := len(titles), 0
	for i := 0; i < titlesLastIdn; i++ {
		if titles[i] == ' ' {
			fmt.Println(i, j)
			if i-j > 2 {
				titles[j] = byteToUpper(titles[j])
			}
			j = i + 1
		} else {
			titles[i] = byteToLower(titles[i])
		}
	}

	return string(titles[:titlesLastIdn-1])
}

func byteToUpper(i byte) byte {
	if i >= 'a' && i <= 'z' {
		return i - 'a' + 'A'
	}

	return i
}

func byteToLower(i byte) byte {
	if i >= 'A' && i <= 'Z' {
		return i - 'A' + 'a'
	}

	return i
}

func test01() {
	s := "capiTalIze tHe titLe"
	successResult := "Capitalize The Title"
	fmt.Println("test01 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", capitalizeTitle(s))
	fmt.Println()
}

func test02() {
	s := "First leTTeR of EACH Word"
	successResult := "First Letter of Each Word"
	fmt.Println("test02 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", capitalizeTitle(s))
	fmt.Println()
}

func test03() {
	s := "i lOve leetcode"
	successResult := "i Love Leetcode"
	fmt.Println("test03 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", capitalizeTitle(s))
	fmt.Println()
}

func test04() {
	s := "L hV"
	successResult := "l hv"
	fmt.Println("test04 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", capitalizeTitle(s))
	fmt.Println()
}

func test05() {
	s := "ZW Cl pyR uoC"
	successResult := "zw cl Pyr Uoc"
	fmt.Println("test05 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", capitalizeTitle(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
