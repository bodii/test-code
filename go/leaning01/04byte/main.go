package main

import (
	"fmt"
	"math"
)

func main() {
	s := "Hello, World"

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}


	s2 := "hello"
	// rune 把字符串强制换成一个rune切片
	s3 := []rune(s2)
	s3[0] = 'b'
	fmt.Println(string(s3))

	c1 := "北"
	c2 := '北'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := 'H'
	fmt.Printf("%d\n", c3)

	c4 := byte('H')
	fmt.Printf("%d\n", c4)

	sqrtDemo()

	n := 10
	var f float64
	f = float64(n)
	fmt.Printf("f: %f\n", f)

	c5 := "hello中国人"
	c6 := []rune(c5)

	l := 0
	for i:=0; i < len(c6); i ++ {
		if c6[i] <= 40869 && c6[i] >= 19968 {
			l ++
		}
	}

	fmt.Printf("%s chinese len: %d\n", c5, l)
}


func sqrtDemo() {
	var a, b = 3, 4
	var c int

	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Printf("c: %d\n", c)
}
