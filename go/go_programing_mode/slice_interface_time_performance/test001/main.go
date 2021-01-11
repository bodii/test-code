package main

import (
	"bytes"
	"fmt"
)

func test01() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99

	fmt.Printf("foo: %v, bar: %v\n", foo, bar)
}

func test02() {
	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)
	a[2] = 42

	fmt.Printf("a: %v, b: %v\n", a, b)
}

func test03() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	// dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))
	/*
		dir1 => AAAA
		dir2 => BBBBBBBBB
		dir1 => AAAAsuffix
		dir2 => uffixBBBB
	*/
}

func main() {
	// test01()
	// test02()
	test03()
}
