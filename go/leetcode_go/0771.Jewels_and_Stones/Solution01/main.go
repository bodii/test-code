package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	nums := 0
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(jewels); j++ {
			if stones[i] == jewels[j] {
				nums++
			}
		}
	}

	return nums
}

func test01() {
	J, S := "aA", "aAAbbbb"
	fmt.Println("J:", J, "S:", S)
	fmt.Println("success result: [", 3, "]")
	fmt.Println(numJewelsInStones(J, S))
}

func test02() {
	J, S := "z", "ZZ"
	fmt.Println("J:", J, "S:", S)
	fmt.Println("success result: [", 0, "]")
	fmt.Println(numJewelsInStones(J, S))
}

func main() {
	test01()
	test02()
}
