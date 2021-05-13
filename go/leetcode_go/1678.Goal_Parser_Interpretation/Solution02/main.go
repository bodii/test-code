package main

import (
	"fmt"
)

func interpret(command string) string {
	res := make([]byte, 0)
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res = append(res, 'G')
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				res = append(res, 'o')
				i++
			} else {
				res = append(res, 'a', 'l')
				i += 3
			}
		}
	}

	return string(res)
}

func test01() {
	command := "G()(al)"
	succResult := "Goal"
	fmt.Println("test01 command:=", command)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", interpret(command))
	fmt.Println()
}

func test02() {
	command := "G()()()()(al)"
	succResult := "Gooooal"
	fmt.Println("test02 command:=", command)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", interpret(command))
	fmt.Println()
}

func test03() {
	command := "(al)G(al)()()G"
	succResult := "alGalooG"
	fmt.Println("test03 command:=", command)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", interpret(command))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
