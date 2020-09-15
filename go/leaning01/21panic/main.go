package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	// recover() 必须搭配defer使用
	// defer 必须在之前使用
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接...")
	}()

	panic("painc in func B")

	fmt.Println("func B")
}


func funcC() {
	fmt.Println("func c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
