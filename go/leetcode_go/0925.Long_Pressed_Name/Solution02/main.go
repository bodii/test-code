package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}

	nameLen := len(name)
	typedLen := len(typed)

	if nameLen > typedLen || (nameLen == typedLen && name != typed) {
		return false
	}

	j := 0
	for i := 0; i < typedLen; i++ {
		if j < nameLen && name[j] == typed[i] {
			j++
		} else if i > 0 && typed[i] == typed[i-1] {
		} else {
			return false
		}
	}

	return j == nameLen
}

func test01() {
	name := "alex"
	typed := "aaleex"
	success := true

	fmt.Println("test01 call func isLongPressedName param name:=", name, " typed:=", typed)
	fmt.Println("success result:", success)
	fmt.Println("result:", isLongPressedName(name, typed))
	fmt.Println()
}

func test02() {
	name := "saeed"
	typed := "ssaaedd"
	success := false

	fmt.Println("test02 call func isLongPressedName param name:=", name, " typed:=", typed)
	fmt.Println("success result:", success)
	fmt.Println("result:", isLongPressedName(name, typed))
	fmt.Println()
}

func test03() {
	name := "leelee"
	typed := "lleeelee"
	success := true

	fmt.Println("test03 call func isLongPressedName param name:=", name, " typed:=", typed)
	fmt.Println("success result:", success)
	fmt.Println("result:", isLongPressedName(name, typed))
	fmt.Println()
}

func test04() {
	name := "laiden"
	typed := "laiden"
	success := true

	fmt.Println("test04 call func isLongPressedName param name:=", name, " typed:=", typed)
	fmt.Println("success result:", success)
	fmt.Println("result:", isLongPressedName(name, typed))
	fmt.Println()
}

func test05() {
	name := "laiden"
	typed := "laidene"
	success := false

	fmt.Println("test05 call func isLongPressedName param name:=", name, " typed:=", typed)
	fmt.Println("success result:", success)
	fmt.Println("result:", isLongPressedName(name, typed))
	fmt.Println()
}

func test06() {
	name := "laiden"
	typed := "laidenn"
	success := true

	fmt.Println("test06 call func isLongPressedName param name:=", name, " typed:=", typed)
	fmt.Println("success result:", success)
	fmt.Println("result:", isLongPressedName(name, typed))
	fmt.Println()
}

func test07() {
	name := "vtkgn"
	typed := "vttkgnn"
	success := true

	fmt.Println("test07 call func isLongPressedName param name:=", name, " typed:=", typed)
	fmt.Println("success result:", success)
	fmt.Println("result:", isLongPressedName(name, typed))
	fmt.Println()
}

func test08() {
	name := "pyplrz"
	typed := "ppyypllr"
	success := false

	fmt.Println("test08 call func isLongPressedName param name:=", name, " typed:=", typed)
	fmt.Println("success result:", success)
	fmt.Println("result:", isLongPressedName(name, typed))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
}
