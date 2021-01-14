package main

import (
	"fmt"
	"reflect"
)

// Transform ...
func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

// TransformPlace ...
func TransformPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ")")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Cap())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}

	return sliceOutType.Interface()
}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}

	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}

	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}

	return true
}

func test01() {
	list := []string{"1", "2", "3", "4", "5"}
	result := Transform(list, func(a string) string {
		return a + a + a
	})

	fmt.Printf("%v\n", result)
}

func test02() {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	TransformPlace(list, func(a int) int {
		return a * 3
	})

	fmt.Printf("%v\n", list)
}

// Employee struct
type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

func test03() {
	list := []Employee{
		{"Jack", 44, 0, 8000},
		{"Tom", 34, 10, 5000},
		{"Bob", 23, 5, 9000},
		{"Alice", 26, 0, 4000},
	}

	fmt.Printf("%v\n", list)

	result := TransformPlace(list, func(e Employee) Employee {
		e.Salary += 100
		e.Age++
		return e
	})

	fmt.Printf("%v\n %v\n", result, list)
}

func main() {
	test01()

	test02()

	test03()
}
