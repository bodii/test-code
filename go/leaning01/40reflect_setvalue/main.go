package main

import (
	"fmt"
	"reflect"
)

/*
	想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改
	变量值。
	而反射中使用专有的Elem()方法来获取指针对应的值。
*/

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a int64 = 100

	// reflectSetValue1(a)
	// panic: reflect: reflect.Value.SetInt using unaddressable value

	reflectSetValue2(&a)

	fmt.Println(a)
}
