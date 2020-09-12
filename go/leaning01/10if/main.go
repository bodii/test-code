package main

import (
	"fmt"
)

func main() {
	age := 19
	if age >= 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}

	if age > 35 {
		fmt.Println("大于35岁")
	} else if age > 18 {
		fmt.Println("大于18岁")
	} else {
		fmt.Println("小于等于18岁")
	}


	if anNum := 19; anNum > 18 {
		fmt.Println("大于18")
	}
}
