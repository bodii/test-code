package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Map generic
func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}

	return result
}

func test01() {
	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}
	squaredArr := Map(nums, square)
	fmt.Println(squaredArr)
}

func test02() {
	upcase := func(s string) string {
		return strings.ToUpper(s)
	}
	strs := []string{"Jack", "Tom", "Alice"}
	upstrs := Map(strs, upcase)
	fmt.Println(upstrs)
}

func main() {
	test01()

	test02()
}
