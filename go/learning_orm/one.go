package main

import (
	"fmt"
	"reflect"
)

// Account type struct
type Account struct {
	Username string
	Password string
}

func main() {
	typ := reflect.Indirect(reflect.ValueOf(&Account{})).Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Println(field.Name)
	}
}
