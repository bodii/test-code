package main

import (
	"fmt"
)

func permutation(S string) []string {
	result := make([]string, 0)

	var dfs func(*[]string, []byte, int)
	dfs = func(res *[]string, s []byte, index int) {
		if index == len(s) {
			*res = append(*res, string(s))
		}

		for i := index; i < len(s); i++ {
			s[i], s[index] = s[index], s[i]
			dfs(res, s, index+1)
			s[i], s[index] = s[index], s[i]
		}
	}

	dfs(&result, []byte(S), 0)

	return result
}

func test01() {
	S := "qwe"
	successResult := []string{"qwe", "qew", "wqe", "weq", "ewq", "eqw"}
	fmt.Println("test01 s:=", S)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", permutation(S))
	fmt.Println()
}

func test02() {
	S := "ab"
	successResult := []string{"ab", "ba"}
	fmt.Println("test01 s:=", S)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", permutation(S))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
