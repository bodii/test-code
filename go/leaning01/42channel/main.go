package main

import "fmt"

/*
	Go语言的并发模型是CSP（Communicating Sequential Processes），
	提倡 通过通信共享内存 而不是通过共享内存而实现通信。

	Go语言中的通信（channel）是一种特殊的类型。通道像一个传送带或者队列，
	总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，
	也就是声明channel的时候需要为其指定元素类型。
*/

var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道

/*
	无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值
	才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine
	将阴塞，直到另一个goroutine在该通道上发送一个值。

	使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
*/
func recv(c chan int) {
	ret := <-c
	fmt.Println("received success", ret)
}

func makeChanDemo() {
	ch4 := make(chan int)
	ch5 := make(chan bool)
	ch6 := make(chan []int)
	fmt.Println(ch4, ch5, ch6)

	// defer close(ch1)

	// // 将一个值发送到通道中
	// ch1 <- 10

	// // 从一个通道中接收值
	// x := <-ch1
	// fmt.Println(x)

	go recv(ch4) // 启用goroutine从通道接收值
	ch4 <- 10
	fmt.Println("sent success")
}

/*
	只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素
	的数量。
	可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通过的容量。
*/
func makeChanDemo2() {
	// 创建一个容量为1的有缓冲区通道
	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("sent success")
}

func makeChanDemo3() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0～100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印,当通道被关闭时就会退出
	for i := range ch2 {
		fmt.Println(i)
	}
}

func main() {
	makeChanDemo()

	makeChanDemo2()

	makeChanDemo3()
}
