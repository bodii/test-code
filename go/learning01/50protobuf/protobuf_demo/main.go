package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bodii/test-code/go/learning01/50protobuf/protobuf_demo/address"
)

func protobufDemo() {
	var cb address.ContactBook

	p1 := &address.Person{
		Name:   "jack",
		Gender: address.GenderType_MALE,
		Number: "123456789",
	}

	fmt.Println(p1)
	cb.Persons = append(cb.Persons, p1)
	// serialized
	var data []byte
	data, err := p1.XXX_Marshal(data, true)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("pserson id:%d name:%s gender:%s number:%s\n",
		p1.GetId(), p1.GetName(), p1.GetGender(), p1.GetNumber())

	path, _ := os.Getwd()
	path += "/50protobuf/protobuf_demo/"

	ioutil.WriteFile(path+"proto.data", data, 0644)

	data2, err := ioutil.ReadFile(path + "proto.data")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}

	var p2 address.Person
	p2.XXX_Unmarshal(data2)
	fmt.Println(p2)
}

func main() {
	protobufDemo()
}
