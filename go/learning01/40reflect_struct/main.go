package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study() string {
	msg := "study hard, improve every day"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "sleep well and grow up quickly"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func printField(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v, json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

func main() {
	stu1 := student{
		Name:  "jack",
		Score: 56,
	}

	printMethod(stu1)

	printField(stu1)
}
