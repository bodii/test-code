package main

import "fmt"

func example01() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ':
		case '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))
}

func example02() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ':
			fallthrough
		case '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))
}

func example03() {
	// fallthrough 不会判断下一条case的expr结果是否为true,并继续执行下一条内容
	switch {
	case false:
		fmt.Println("the integer was <= 4")
		fallthrough
	case true:
		fmt.Println("The integer was <= 5")
		fallthrough
	case false:
		fmt.Println("The integer was <= 6")
		fallthrough
	case true:
		fmt.Println("The integer was <= 7")
		fallthrough
	case false:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default")
		// fallthrough // cannot fallthrough final case in switch
	}
}

func main() {
	example01()
	example02()
	example03()
}
