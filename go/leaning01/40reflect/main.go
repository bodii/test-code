package main

import (
	"fmt"
	"reflect"
)

/*
	go 语言的反射中像数组、切片、Map、指针等类型的变量，它们的.name()都是返回空。
*/
type myInt int64

type person struct {
	name string
	age  int
}

type book struct{ title string }

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a)

	var b int64 = 100
	reflectType(b)

	var c myInt
	reflectType(c)

	var d rune
	reflectType(d)

	p := person{
		name: "jack",
		age:  18,
	}

	e := book{title: "golang learning"}

	reflectType(p)
	reflectType(e)

	reflectValue(a)
	reflectValue(b)

	f := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", f)
}
