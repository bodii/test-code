package main

import "fmt"

func example01() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)

	in = data
	fmt.Println(in, in == nil)
}

func example02() {
	doIt := func(arg int) interface{} {
		var result *struct{} = nil

		if arg > 0 {
			result = &struct{}{}
		}

		return result
	}

	if res := doIt(-1); res != nil {
		fmt.Println("Good result: ", res)
	} else {
		fmt.Println("Bad result: ", res)
	}
}

func main() {
	example01()

	example02()
}
