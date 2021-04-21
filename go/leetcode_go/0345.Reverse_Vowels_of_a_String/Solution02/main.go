package main

import "fmt"

func reverseVowels(s string) string {
	len := len(s)
	end := len
	n := []byte(s)
	for i := 0; i < len; i++ {
		if i >= end {
			break
		}
		if isVowels(n[i]) {
			for j := end - 1; j > i; j-- {
				if isVowels(n[j]) {
					// n[i], n[j] = n[j], n[i]
					swap(&n, i, j)
					end = j
					break
				}
			}
		}
	}

	return string(n)
}

func swap(n *[]byte, i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func isVowels(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func test01() {
	s := "hello"
	fmt.Println("func test01 for result: ", reverseVowels(s))
}

func test02() {
	s := "leetcode"
	fmt.Println("func test02 for result: ", reverseVowels(s))
}

func test03() {
	s := "aA"
	fmt.Println("func test03 for result: ", reverseVowels(s))
}

func main() {
	test01()
	test02()
	test03()
}
