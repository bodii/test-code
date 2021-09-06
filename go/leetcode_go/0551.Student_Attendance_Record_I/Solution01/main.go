package main

import (
	"fmt"
)

func checkRecord(s string) bool {
	absents, lates := 0, 0
	for _, v := range s {
		if v == 'A' {
			absents++
			if absents >= 2 {
				return false
			}
			lates = 0
		} else if v == 'L' {
			lates++
			if lates >= 3 {
				return false
			}
		} else {
			lates = 0
		}
	}

	return true
}

func test01() {
	s := "PPALLP"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkRecord(s))
	fmt.Println()
}

func test02() {
	s := "PPALLL"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkRecord(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
