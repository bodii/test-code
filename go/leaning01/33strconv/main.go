package main

import (
	"fmt"
	"strconv"
)

func atoiDemo() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
		return
	}

	fmt.Printf("type:%T value:%#v\n", i1, i1)
}

func itoaDemo() {
	i := 200
	s2 := strconv.Itoa(i)
	fmt.Printf("type:%T value:%#v\n", s2, s2)
}

func  parseBoolDemo() {
	i := "t"
	b, err := strconv.ParseBool(i)
	if err != nil {
		fmt.Println("string type conv boolean type failed! err:", err)
		return 
	}
	fmt.Printf("type:%T value:%#v\n", b, b)
}

func parseIntDemo() {
	i := "100"
	pi, err := strconv.ParseInt(i, 8, 8)
	if err != nil {
		fmt.Println("conv failed! err:", err)
		return
	}

	fmt.Printf("type:%T value:%#v\n", pi, pi)
}

func formatBool() {
	t := false
	f := strconv.FormatBool(t)
	fmt.Println(f)
}

func formatFloat() {
	f := 2.454543
	fmt.Println(strconv.FormatFloat(f, 'E', -1, 64))
}


func isPrintDemo() {
	s := ' '
	fmt.Println(strconv.IsPrint(s))
}

func canBackQuoteDemo() {
	s := "kdsfldsf\ndsfsdf\t dsfdf1243"
	fmt.Println(strconv.CanBackquote(s))

	s2 := "kdsfldsfdsfsdf dsfdf1243"
	fmt.Println(strconv.CanBackquote(s2))
}


func main() {
	atoiDemo()
	itoaDemo()
	parseBoolDemo()
	parseIntDemo()
	formatBool()
	formatFloat()
	isPrintDemo()
	canBackQuoteDemo()
}
