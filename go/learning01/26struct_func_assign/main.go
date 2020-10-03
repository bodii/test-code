package main

import "fmt"


// 结构体
type dog struct {
	name string
}


// 构造函数
func newDog(name string) dog {
	return dog {
		name: name,
	}
}

// 方法是作用于特定类型的函数
func (d dog) shout() {
	fmt.Println(d.name, "woowoowoo!!!")
}

func main() {
	d := newDog("xiaohe")
	d.shout()	
}
