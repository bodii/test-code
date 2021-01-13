package main

import (
	"fmt"
	"reflect"
)

// Reduce ...
func Reduce(slice, pairFunc, zero interface{}) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}

	len := sliceInType.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sliceInType.Index(0)
	}

	elemType := sliceInType.Type().Elem()
	fn := reflect.ValueOf(pairFunc)
	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func(" + t + ", " + t + ") " + t)
	}

	var ins [2]reflect.Value
	ins[0] = sliceInType.Index(0)
	ins[1] = sliceInType.Index(1)
	out := fn.Call(ins[:])[0]

	for i := 2; i < len; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}

	return out.Interface()
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

func main() {
	list := []string{"Jack", "Tome", "Lias"}
	zero := "default"

	z := Reduce(list, func(s1, s2 string) string {
		return s1 + ", " + s2
	}, zero)
	fmt.Printf("%v\n", z)
}
