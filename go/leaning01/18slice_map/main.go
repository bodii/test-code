package main

import (
	"fmt"
)

func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "name vlaue"
	mapSlice[0]["age"] = "18 year"
	mapSlice[0]["sex"] = "other"

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
