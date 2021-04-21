package main

import "fmt"

func reverseVowels(s string) string {
	len := len(s)
	n := []byte(s)
	end := len
	for i := 0; i < len; i++ {
		if n[i] == 'a' || n[i] == 'e' || n[i] == 'i' || n[i] == 'o' || n[i] == 'u' ||
			n[i] == 'A' || n[i] == 'E' || n[i] == 'I' || n[i] == 'O' || n[i] == 'U' {
			for j := end - 1; j > i; j-- {
				if n[j] == 'a' || n[j] == 'e' || n[j] == 'i' || n[j] == 'o' || n[j] == 'u' ||
					n[j] == 'A' || n[j] == 'E' || n[j] == 'I' || n[j] == 'O' || n[j] == 'U' {
					n[i], n[j] = n[j], n[i]
					end = j
					break
				}
			}
		}
	}

	return string(n)
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
