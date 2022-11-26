package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			// 如果once.Do(f)被多次调用，只有第一次调用会执行f
			// 即使f每次调用Do提供的f值不同。需要给每个要执行仅一次的函数都建立一个Once类型的实例
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		d := <-done
		fmt.Println(d)
	}

	// output: Only once
}
