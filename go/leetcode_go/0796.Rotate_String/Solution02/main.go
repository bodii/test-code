package main

import (
	"fmt"
)

func rotateString(A string, B string) bool {
	ALen := len(A)

	if (ALen == 0 && A == B) || A == B {
		return true
	}

	for i := 0; i < ALen; i++ {
		if B[0] == A[i] {
			if A[i:]+A[0:i] == B {
				return true
			}
		}
	}

	return false
}

func test01() {
	A, B := "abcde", "cdeab"
	success := true

	fmt.Println("test01 call func rotateString param A:=", A, " B:=", B)
	fmt.Println("success result:", success)
	fmt.Println("result:", rotateString(A, B))
	fmt.Println()
}

func test02() {
	A, B := "abcde", "abced"
	success := false

	fmt.Println("test02 call func rotateString param A:=", A, " B:=", B)
	fmt.Println("success result:", success)
	fmt.Println("result:", rotateString(A, B))
	fmt.Println()
}

func test03() {
	A, B := "eckadf", "kadfec"
	success := true

	fmt.Println("test03 call func rotateString param A:=", A, " B:=", B)
	fmt.Println("success result:", success)
	fmt.Println("result:", rotateString(A, B))
	fmt.Println()
}

func test04() {
	A, B := "bbbacddceeb", "ceebbbbacdd"
	success := true
	fmt.Println("test04 call func rotateString param A:=", A, " B:=", B)
	fmt.Println("success result:", success)
	fmt.Println("result:", rotateString(A, B))
	fmt.Println()

}

func main() {
	test01()
	test02()
	test03()
	test04()
}
