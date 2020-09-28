package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // gorountine结束就登记
	fmt.Println("Hello Goroutine!", i)
}

/*
	10个goroutine并发执行，而goroutine的调度是随机的
*/
func main() {
	// 启动另外一个goroutine去执行hello函数
	// go hello()
	// fmt.Println("main goroutine done!")
	// time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记
		go hello(i)
	}

	wg.Wait() // 等待所有登记的goroutine都结束
}
