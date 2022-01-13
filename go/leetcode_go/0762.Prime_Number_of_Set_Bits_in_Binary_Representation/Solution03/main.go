package main

import "fmt"

func countPrimeSetBits(left int, right int) int {
	count := 0
	for ; left <= right; left++ {
		switch countOnes(left) {
		case 2, 3, 5, 7, 11, 13, 17, 19:
			count++
		}
	}

	return count
}

func countOnes(i int) int {
	ones := 0
	for ; i > 0; i >>= 1 {
		if i&1 == 1 {
			ones++
		}
	}
	fmt.Println(ones)

	return ones
}

func test01() {
	left, right := 6, 10
	succResult := 4
	fmt.Println("test01:")
	fmt.Println("left:", left, " right:=", right)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", countPrimeSetBits(left, right))
	fmt.Println()
}

func test02() {
	left, right := 10, 15
	succResult := 5
	fmt.Println("test02:")
	fmt.Println("left:", left, " right:=", right)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", countPrimeSetBits(left, right))
	fmt.Println()
}

func test03() {
	left, right := 842, 888
	succResult := 23
	fmt.Println("test03:")
	fmt.Println("left:", left, " right:=", right)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", countPrimeSetBits(left, right))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
