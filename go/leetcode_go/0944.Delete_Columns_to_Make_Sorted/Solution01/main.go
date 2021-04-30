package main

import "fmt"

func minDeletionSize(strs []string) int {
	rest := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			// fmt.Println(string(strs[j-1][i]), string(strs[j][i]))
			if strs[j-1][i] > strs[j][i] {
				rest++
				break
			}
		}
	}
	return rest
}

func test01() {
	strs := []string{"cba", "daf", "ghi"}
	fmt.Printf("test01 %v is 1 result: %d\n", strs, minDeletionSize(strs))
}

func test02() {
	strs := []string{"a", "b"}
	fmt.Printf("test02 %v is 0 result: %d\n", strs, minDeletionSize(strs))
}

func test03() {
	strs := []string{"zyx", "wvu", "tsr"}
	fmt.Printf("test03 %v is 3 result: %d\n", strs, minDeletionSize(strs))
}

func test04() {
	strs := []string{"zy", "wv"}
	fmt.Printf("test04 %v is 2 result: %d\n", strs, minDeletionSize(strs))
}
func test05() {
	strs := []string{"rrjk", "furt", "guzm"}
	fmt.Printf("test05 %v is 2 result: %d\n", strs, minDeletionSize(strs))
}

func test06() {
	strs := []string{"rrj", "fur", "guz", "ktm"}
	fmt.Printf("test06 %v is 3 result: %d\n", strs, minDeletionSize(strs))
}

func test07() {
	strs := []string{"rrjk", "furt", "guzm"}
	fmt.Printf("test07 %v is 2 result: %d\n", strs, minDeletionSize(strs))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
}
