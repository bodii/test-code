package main

import "fmt"

func rotateString(A string, B string) bool {
	len := len(A)
	if len == 0 && A == B {
		return true
	}

	for i := 0; i < len; i++ {
		A = move(A)
		if A == B {
			return true
		}
	}
	return false
}

func move(A string) string {
	A = A[1:] + A[0:1]
	return A
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

func main() {
	test01()
	test02()
}
