package main

import (
	"flag"
	"fmt"
)

var Input_pstrName = flag.String("name", "gerry", "input ur name")
var Input_piAge = flag.Int("age", 20, "input ur age")
var Input_flagvar int

func Init() {
	flag.IntVar(&Input_flagvar, "flagname", 123, "help message for flagname")
}

func main() {
	Init()
	flag.Parse()

	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("name =", *Input_pstrName)
	fmt.Println("age =", *Input_piAge)
	fmt.Println("flagname =", Input_flagvar)
}
