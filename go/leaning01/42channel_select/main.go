package main

import "fmt"

/*
	select 多路利用

	e.g.
	for {
		data, ok := <-ch1
		data, ok := <-ch2
	}
	这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。为了
	应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。

	select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个
	case会对应用一个通道的通信（接收或发送）过程。select会一直等待，直到某个
	case的通信操作完成时，就会执行case分支对应的语句。
	select　{
	case <-ch1:
		...
	case data := <-ch2:
		...
	case ch3 <-data:
		...
	default:
		默认操作
	}

	可处理一个或多个channel的发送接收操作。
	如果多个case同时满足，select会随机选择一个。
	对于没有case的select{}会一直等待，可用于阻塞main函数
*/

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
