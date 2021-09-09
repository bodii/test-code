package main

import (
	"fmt"
)

func isThree(n int) bool {
	count := 0

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if n/i == i {
				count++
			} else {
				count += 2
			}
		}

		if count > 3 {
			break
		}
	}

	return count == 3
}

func test01() {
	n := 2
	succResult := false
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isThree(n))
	fmt.Println()
}

func test02() {
	n := 4
	succResult := true
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isThree(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
