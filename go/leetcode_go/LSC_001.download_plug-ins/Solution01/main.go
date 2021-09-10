package main

import "fmt"

func leastMinutes(n int) int {
	count := 1

	for i := 1; i < n; i *= 2 {
		count++
	}

	return count
}

func test01() {
	n := 2
	succResult := 2
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", leastMinutes(n))
	fmt.Println()
}

func test02() {
	n := 4
	succResult := 3
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", leastMinutes(n))
	fmt.Println()
}

func test03() {
	n := 3
	succResult := 3
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", leastMinutes(n))
	fmt.Println()
}

func test04() {
	n := 1
	succResult := 1
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", leastMinutes(n))
	fmt.Println()
}

func test05() {
	n := 5
	succResult := 4
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", leastMinutes(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
