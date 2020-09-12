package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5, 10)
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int,0, 5)
	fmt.Println(s2, len(s2), cap(s2))

	s3 := s1
	s3[0] = 100
	fmt.Println(s1, s3)

	for i:=0; i< len(s3);i++ {
		fmt.Println(s3[i])
	}

	for k, v := range s3 {
		fmt.Println(k, v)
	}

}
