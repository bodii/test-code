package main

import (
	"fmt"
	"strconv"
)

func insertBits(N int, M int, i int, j int, successResult int) int {
	M <<= i
	for k := i; k <= j; k++ {
		if N&(1<<k) != 0 {
			N ^= 1 << k
		}
	}
	return N | M
}

func test01() {
	N, M := 1024, 19
	i, j := 2, 6
	successResult := 1100
	fmt.Println("test01 N:=", N, " M:=", M)
	fmt.Println("i:=", i, " j:=", j)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", insertBits(N, M, i, j, successResult))
	fmt.Println()
}

func test02() {
	N, M := 0, 31
	i, j := 0, 4
	successResult := 31
	fmt.Println("test02 N:=", N, " M:=", M)
	fmt.Println("i:=", i, " j:=", j)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", insertBits(N, M, i, j, successResult))
	fmt.Println()
}

func test03() {
	N, M := 10, 31
	i, j := 2, 6
	successResult := 126
	fmt.Println("test03 N:=", N, " M:=", M)
	fmt.Println("i:=", i, " j:=", j)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", insertBits(N, M, i, j, successResult))
	fmt.Println()
}

func test04() {
	N, M := 3, 31
	i, j := 1, 5
	successResult := 63
	fmt.Println("test04 N:=", N, " M:=", M)
	fmt.Println("i:=", i, " j:=", j)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", insertBits(N, M, i, j, successResult))
	fmt.Println()
}

func test05() {
	N, M := 3, 31
	i, j := 3, 5
	successResult := 251
	fmt.Println("test05 N:=", N, " M:=", M)
	fmt.Println("i:=", i, " j:=", j)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", insertBits(N, M, i, j, successResult))
	fmt.Println()
}

func test06() {
	N, M := 1143207437, 1017033
	i, j := 11, 31
	successResult := 2082885133
	fmt.Println("test06 N:=", N, " M:=", M)
	fmt.Println("i:=", i, " j:=", j)
	fmt.Println("success result:", successResult)
	fmt.Println("success binary:", strconv.FormatInt(int64(successResult), 2))
	fmt.Println("result: ", insertBits(N, M, i, j, successResult))
	fmt.Println()
}

func test07() {
	N, M := 2032243561, 10
	i, j := 24, 29
	successResult := 1243714409
	fmt.Println("test07 N:=", N, " M:=", M)
	fmt.Println("i:=", i, " j:=", j)
	fmt.Println("success result:", successResult)
	fmt.Println("success binary:", strconv.FormatInt(int64(successResult), 2))
	fmt.Println("result: ", insertBits(N, M, i, j, successResult))
	fmt.Println()

}

func test08() {
	N, M := 482311233, 0
	i, j := 21, 22
	successResult := 480214081
	fmt.Println("test08 N:=", N, " M:=", M)
	fmt.Println("i:=", i, " j:=", j)
	fmt.Println("success result:", successResult)
	fmt.Println("success binary:", strconv.FormatInt(int64(successResult), 2))
	fmt.Println("        binary:", "11010001000011001001101101001")
	fmt.Println("result: ", insertBits(N, M, i, j, successResult))
	fmt.Println()
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
