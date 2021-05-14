package main

import "fmt"

func numberOfMatches(n int) int {
	pairs := 0

	for n > 1 {
		pair := n / 2
		if !isEven(n) {
			n = pair + 1
		} else {
			n = pair
		}
		// fmt.Println(pair, n)
		pairs += pair
	}

	return pairs
}

func isEven(a int) bool {
	return a%2 == 0
}

func test01() {
	n := 7
	succResult := 6
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numberOfMatches(n))
	fmt.Println()
}

func test02() {
	n := 14
	succResult := 13
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numberOfMatches(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
