package main

import "fmt"

func numWays(n int, relation [][]int, k int) int {
	res := make([]int, n)
	res[0] = 1

	for i := 0; i < k; i++ {
		tmp := make([]int, n)
		for _, v := range relation {

			tmp[v[1]] += res[v[0]]
			fmt.Println(v)
		}
		res = tmp
	}

	return res[n-1]
}

func test01() {
	n := 5
	relation := [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}
	k := 3
	succResult := 3

	fmt.Println("test01 n:=", n, " relation:=", relation, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numWays(n, relation, k))
	fmt.Println()
}

func test02() {
	n := 3
	relation := [][]int{{0, 2}, {2, 1}}
	k := 2
	succResult := 0

	fmt.Println("test01 n:=", n, " relation:=", relation, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numWays(n, relation, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
