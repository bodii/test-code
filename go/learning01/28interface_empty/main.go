package main

import (
	"fmt"
)

// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("%#v %v\n", a, a)
}

// 空接口
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "jack"
	m1["age"] = 18
	m1["merried"] = true
	m1["hobby"] = [...]string{"song", "dancing", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
