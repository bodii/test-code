package main

import (
	"fmt"
)

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ruleKeys := map[string]int{"type": 0, "color": 1, "name": 2}
	nums := 0
	for _, v := range items {
		if v[ruleKeys[ruleKey]] == ruleValue {
			nums++
		}
	}

	return nums
}

func test01() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	succResult := 1
	fmt.Println("test01 item:=", items)
	fmt.Println("ruleKey:=", ruleKey, " ruleVlaue:=", ruleValue)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countMatches(items, ruleKey, ruleValue))
	fmt.Println()
}

func test02() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}
	ruleKey := "type"
	ruleValue := "phone"
	succResult := 2
	fmt.Println("test01 item:=", items)
	fmt.Println("ruleKey:=", ruleKey, " ruleVlaue:=", ruleValue)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countMatches(items, ruleKey, ruleValue))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
