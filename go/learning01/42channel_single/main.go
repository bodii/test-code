package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

/*
	chan<- int 是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行
	接收操作；
	<-chan int 是一个只读单向通道(只能从其读取int类型值)，可以对其执行接收操作但是不能执行发
	送操作。

	在函数传参及任何赋值操作中， 可以将双向通道转换为单向通道，但反过来是不可以的。

	关闭已经关闭的channel也会引发panic
*/
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}

	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
