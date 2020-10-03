package main

import (
	"fmt"
)

// 类型断言
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("the string type is error")
		return
	}
	fmt.Println("伟进来的是一个字符串", str)
}


// 用swith语句作类型判断
func justIfyType(x interface{}) {
	switch v := x.(type) {
		case string:
			fmt.Printf("x is a string, value is %v\n", v)
		case int:
			fmt.Printf("x is a int, value is %v\n", v)
		case bool:
			fmt.Printf("x is a bool, value is %v\n", v)
		default:
			fmt.Println("unsupport type!")
	}
}

/*
关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式
进行处理时才需要定义接口，不要为了接口而写接口，那样只会增加不必要的抽象，
导致不必要的运行时损耗。
*/


func main() {
	assign(100)

	justIfyType(1)
	justIfyType(true)
	justIfyType("a")
	justIfyType('a')
}
