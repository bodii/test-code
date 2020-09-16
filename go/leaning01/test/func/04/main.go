package main

import (
	"fmt"
)

/*
n个台阶，一次可以走1步，也可以走2步，有多少种走法。
*/

func f(n int) int {
	if n == 1 {
		return 1
	} else if n == 2{
		return 2
	}

	return f(n - 1) + f(n - 2)
}

func main() {
	fmt.Println(f(10))	
}
