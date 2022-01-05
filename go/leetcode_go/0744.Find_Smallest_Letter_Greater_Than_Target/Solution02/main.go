package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, v := range letters {
		if v > target {
			return v
		}
	}

	return letters[0]
}

func test01() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')

	succResult := 'c'
	fmt.Println("test01:")
	fmt.Println("letters:=", letters, " target:=", target)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", nextGreatestLetter(letters, target))
	fmt.Println()
}

func test02() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('c')

	succResult := byte('f')
	fmt.Println("test02:")
	fmt.Println("letters:=", letters, " target:=", target)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", nextGreatestLetter(letters, target))
	fmt.Println()
}

func test03() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('d')

	succResult := byte('f')
	fmt.Println("test03:")
	fmt.Println("letters:=", letters, " target:=", target)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", nextGreatestLetter(letters, target))
	fmt.Println()
}

func test04() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('g')

	succResult := byte('j')
	fmt.Println("test04:")
	fmt.Println("letters:=", letters, " target:=", target)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", nextGreatestLetter(letters, target))
	fmt.Println()
}

func test05() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('j')

	succResult := byte('c')
	fmt.Println("test05:")
	fmt.Println("letters:=", letters, " target:=", target)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", nextGreatestLetter(letters, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
