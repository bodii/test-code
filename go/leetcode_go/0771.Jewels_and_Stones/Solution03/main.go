package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	nums := 0
	jewelsMap := make(map[rune]int)

	for _, v := range jewels {
		jewelsMap[v]++
	}

	for _, v := range stones {
		nums += jewelsMap[v]
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
