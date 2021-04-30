package main

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
	s := []byte(strconv.Itoa(num))
	for i := 0; i < len(s); i++ {
		if s[i] == 54 {
			s[i] = 57
			break
		}
	}

	result, _ := strconv.Atoi(string(s))
	return result
}

func test01() {
	num := 9669
	fmt.Println("test01 is 9669 result 9969:", maximum69Number(num))
}

func test02() {
	num := 9996
	fmt.Println("test01 is 9996 result 9999:", maximum69Number(num))
}

func test03() {
	num := 9999
	fmt.Println("test01 is 9999 result 9999:", maximum69Number(num))
}

func main() {
	test01()
	test02()
	test03()
}
