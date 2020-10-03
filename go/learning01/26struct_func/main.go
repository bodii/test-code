package main

import "fmt"

type person struct {
	name string
	age int
}

type dog struct {
	name string
}

// 构造函数
/*
func newPerson(name string, age int) person {
	return person {
		name: name,
		age : age,
	}
}
*/
// 当结构体比较大时，尽量使用结构体指针，减少程序内存的开销
func newPerson(name string, age int) *person {
	return &person {
		name: name,
		age : age,
	}
}


// 结构体返回是值就是拷贝
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p := newPerson("jack", 18)
	fmt.Printf("%#v\n", p)

	d := newDog("xiaohe")
	fmt.Printf("%#v\n", d)
}
