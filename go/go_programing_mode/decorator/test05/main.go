package main

import (
	"fmt"
	"reflect"
)

// Decorator generic decorator
func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}


func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}

type MyFoo func(int, int, int) int
var myFoo MyFoo

func main() {
	Decorator(&myFoo, foo)
	myFoo(1, 2, 3)

	myBar := bar
	Decorator(&myBar, bar)
	myBar("hello,", "world!")
}
