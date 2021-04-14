package main

import "fmt"

func canWinNim(n int) bool {
	return n%4 != 0
}

func test01() {
	fmt.Println("test 01 func the result is : ", canWinNim(5))
}

func main() {
	test01()
}
