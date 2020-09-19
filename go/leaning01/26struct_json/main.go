package main

import (
	"fmt"
	"encoding/json"
)

// struct and json
// 1. 序列化：把go语言中的结构体变量 -> json格式的字符串
// 2. 反序列化: 把json格式的字符串 -> go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}


func main() {
	p := person{
		Name: "jack",
		Age: 18,
	}

	// 序列化
	j, err := json.Marshal(p)
	if err == nil {
		fmt.Printf("marshal: %#v\n", string(j))
	}

	// 反序列化
	str := `{"name":"Tom","age":19}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
