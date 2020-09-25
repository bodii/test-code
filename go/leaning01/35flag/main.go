package main

import (
	"fmt"
	//"os"
	"flag"
	"time"
)


func main() {
	/*
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flay.Duration("duration", 0, "时间间隔")
	*/
	/*
	var (
		name string
		age int
		married bool
		delay time.Duration
	)
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "duration", 0, "时间间隔")
	*/
	/*
	fmt.Printf("%#v\n", flag.Args())
	fmt.Printf("%#v\n", flag.NArg())
	fmt.Printf("%#v\n", flag.NFlag())
	*/

	var (
		name string
		age int
		married bool
		delay time.Duration
	)

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "duration", 0, "时间间隔")

	// 解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
