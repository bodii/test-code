package main

import (
	"fmt"
)

func main() {
	var m = make(map[string]int, 10)
	m["str"] = 1
	m["str2"] = 2
	fmt.Println(m)

	v, ok := m["str"]
	if !ok {
		fmt.Println("'m' is the map type, the key is 'str' is  empty ")
	} else {
		fmt.Println(v)
	}

	// 如果key不存在,值默认为0
	fmt.Println(m["str3"])
	
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	delete(m, "str")
	fmt.Println(m)
	// 删除一个不存在的key
	delete(m, "str3")
	fmt.Println(m)



}
