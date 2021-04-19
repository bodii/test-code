package main

import "fmt"

func failure_example01() {
	recover()
	panic("not good")
	recover()
	fmt.Println("ok")
}

func failure_example02() {
	defer func() {
		func() { recover() }()
	}()
	panic(1)
}

func success_example() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()
	panic("not good")
}

func main() {
	// failure_example01()
	// failure_example02()
	success_example()
}
