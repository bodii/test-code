package main

import (
	"fmt"
	"strings"
)

// MapStrToStr ...
func MapStrToStr(arr []string, fn func(s string) string) []string {
	newArray := []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}

	return newArray
}

// MapStrToInt ...
func MapStrToInt(arr []string, fn func(s string) int) []int {
	newArray := []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}

	return newArray
}

func main() {
	list := []string{"Jack", "Tome", "Lias"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y)
}
