package main

import (
	"fmt"
)

func findRestaurant(list1 []string, list2 []string) []string {
	list1Len, list2Len := len(list1), len(list2)

	listMap := make(map[string]int)
	min := -1
	var result []string

	for i := 0; i < list1Len; i++ {
		listMap[list1[i]] = i
	}

	for i := 0; i < list2Len; i++ {
		if _, ok := listMap[list2[i]]; !ok {
			continue
		}

		if min == -1 || listMap[list2[i]]+i < min {
			result = append([]string{}, list2[i])
			min = listMap[list2[i]] + i
		} else if listMap[list2[i]]+i == min {
			result = append(result, list2[i])
		}
	}

	return result
}

func test01() {
	list1, list2 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	succResult := []string{"Shogun"}
	fmt.Println("test01 list1:=", list1, " list2:=", list2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findRestaurant(list1, list2))
	fmt.Println()
}

func test02() {
	list1, list2 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"}
	succResult := []string{"Shogun"}
	fmt.Println("test02 list1:=", list1, " list2:=", list2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findRestaurant(list1, list2))
	fmt.Println()
}

func test03() {
	list1, list2 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	succResult := []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	fmt.Println("test03 list1:=", list1, " list2:=", list2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findRestaurant(list1, list2))
	fmt.Println()
}

func test04() {
	list1, list2 := []string{"S", "TEXP", "BK", "KFC"},
		[]string{"KFC", "BK", "S"}
	succResult := []string{"S"}
	fmt.Println("test04 list1:=", list1, " list2:=", list2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findRestaurant(list1, list2))
	fmt.Println()
}

func test05() {
	list1, list2 := []string{"k", "KFC"},
		[]string{"k", "KFC"}
	succResult := []string{"k"}
	fmt.Println("test05 list1:=", list1, " list2:=", list2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findRestaurant(list1, list2))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
