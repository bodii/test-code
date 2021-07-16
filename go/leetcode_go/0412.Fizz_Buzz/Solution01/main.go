package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
			continue
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
			continue
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
			continue
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}

func test01() {
	n := 3
	succResult := []string{"1", "2", "Fizz"}
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fizzBuzz(n))
	fmt.Println()
}

func test02() {
	n := 5
	succResult := []string{"1", "2", "Fizz", "4", "Buzz"}
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fizzBuzz(n))
	fmt.Println()
}

func test03() {
	n := 15
	succResult := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fizzBuzz(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
