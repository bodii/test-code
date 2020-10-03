package main

import "fmt"


// 获取用户输入

func main() {
	var s string
	fmt.Println("input value:")
	fmt.Scan(&s)
	fmt.Println(s)

	var (
		name, class string
		age int
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)

	var (
		name1 string
		age1 int
		class1 string
	)
	fmt.Scanln(&name1, &age1, &class1)
	fmt.Println(name1, age1, class1)
}
