package main

import (
	"fmt"
)

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

var pick int

func guessNumber(n int) int {
	l, r := 0, n
	for l < r {
		t := l + (r-l)/2
		if guess(t) <= 0 {
			r = t
		} else {
			l = t + 1
		}
	}

	return l
}

func test01() {
	n := 10
	pick = 6
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", pick)
	fmt.Println("result:=", guessNumber(n))
	fmt.Println()
}

func test02() {
	n := 1
	pick = 1
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", pick)
	fmt.Println("result:=", guessNumber(n))
	fmt.Println()
}

func test03() {
	n := 2
	pick = 1
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", pick)
	fmt.Println("result:=", guessNumber(n))
	fmt.Println()
}

func test04() {
	n := 2
	pick = 2
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", pick)
	fmt.Println("result:=", guessNumber(n))
	fmt.Println()
}

func test05() {
	n := 5
	pick = 3
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", pick)
	fmt.Println("result:=", guessNumber(n))
	fmt.Println()
}

func test06() {
	n := 3
	pick = 2
	fmt.Println("test06 n:=", n)
	fmt.Println("success result:=", pick)
	fmt.Println("result:=", guessNumber(n))
	fmt.Println()
}

func test07() {
	n := 1000
	pick = 50
	fmt.Println("test07 n:=", n)
	fmt.Println("success result:=", pick)
	fmt.Println("result:=", guessNumber(n))
	fmt.Println()
}

func guess(num int) int {
	fmt.Println(num)
	if pick < num {
		return -1
	} else if pick > num {
		return 1
	}

	return 0
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
}
