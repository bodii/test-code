package main

import (
	"fmt"
	"reflect"
)

func main() {
	// *int type ptr
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())

	// nil value
	fmt.Println("nil IsValid:", reflect.ValueOf(a).IsValid())

	// instantiate
	b := struct{}{}
	fmt.Println("does it exist struct member:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	fmt.Println("does it exist struct method:", reflect.ValueOf(b).MethodByName("abc").IsValid())

	c := map[string]int{}
	fmt.Println("does it  exist  key in map :", reflect.ValueOf(c).MapIndex(reflect.ValueOf("jack")).IsValid())
}
