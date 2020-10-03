package main

import (
	"fmt"
)


func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}


func main() {
	add := adder(100)

	ret := add(200)
	fmt.Println(ret) // 300
	
	ret2 := add(200)
	fmt.Println(ret2) // 500

}
