package main

import (
	"fmt"
	"strings"
)

/*
写一个程序,统计一个字符串中每个单词出现的次数,比如:"how do you do"
中 how=1 do=2 you=1
*/
func main() {
	s := "how do you do"
	arr := strings.Fields(s)
	fmt.Println(arr)

	var m = make(map[string]int, 10)
	for _, v := range arr {
		var _, ok = m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}

	fmt.Println(m)
}
