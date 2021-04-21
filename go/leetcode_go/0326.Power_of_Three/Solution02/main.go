package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPowerOfThree(n int) bool {
	s := strconv.FormatInt(int64(n), 3)
	if '1' == s[0] && strings.ContainsAny(s, "1") {
		return true
	}

	return false
}

func test01() {
	fmt.Println("n = 27 the result:", isPowerOfThree(27))
}

func test02() {
	fmt.Println("n = 3 the result:", isPowerOfThree(3))
}

func test03() {
	fmt.Println("n = -9 is [false] the result:", isPowerOfThree(-9))
}

func test04() {
	fmt.Println("n = 45 is [false] the result:", isPowerOfThree(45))
}

func test05() {
	fmt.Println("n = 5555559 is [false]  the result:", isPowerOfThree(5555559))
}

func test06() {
	fmt.Println("n = 2147483647 is [false] the result:", isPowerOfThree(2147483647))
}

func test07() {
	fmt.Println("n = 11 is [false] the result:", isPowerOfThree(11))
}

func test08() {
	fmt.Println("n = 1 is [true] the result:", isPowerOfThree(1))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
}
