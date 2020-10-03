package main

import "fmt"

// 自定义类型并添加方法
// 不能给别的包里面的类型添加方法，只能给自己的包内的类型添加
type myInt int

func (i myInt) hello() {
	fmt.Println(i, ":我是一个int")
}

func main() {
	i := myInt(8)
	i.hello()
}
