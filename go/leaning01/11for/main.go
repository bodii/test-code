package main

import (
	"fmt"
)


func main() {
	for i:=0; i < 10; i++ {
		fmt.Println(i)
	}

	var i1 uint = 5
	for ; i1 < 8; i1++ {
		fmt.Println(i1)
	}

	i2 := 0
	for ; i2 < 5; {
		fmt.Println(i2)
		i2++
	}

	/*
	for {
		fmt.Println(1)
	}
	*/

	s := "hello world"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
