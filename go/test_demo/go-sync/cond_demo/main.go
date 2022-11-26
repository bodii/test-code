package main

import (
	"fmt"
	"sync"
	"time"
)

var start = false
var done = false

func Soldiers(i int, c *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	c.L.Lock()
	fmt.Printf("大兵%d号等待干饭...\n", i)
	for !start {
		// Wait 解锁lock 同时挂起goroutine
		// 当Wait的goroutine被唤醒的时候，会重新将锁加上
		c.Wait()
	}
	fmt.Printf("大兵%d号开始干饭...\n", i)
	c.L.Unlock()
}

func Waiter(c *sync.Cond) {
	c.L.Lock()

	for !done {
		c.Wait()
	}
	fmt.Println("用餐结束，开始打扫...")

	c.L.Unlock()
}

func Officer(c *sync.Cond) {
	fmt.Println("长官准备中...")
	time.Sleep(5 * time.Second)

	c.L.Lock()
	start = true
	c.L.Unlock()

	// 唤醒的所有等待线程
	c.Broadcast()
}

func CleanUp(c *sync.Cond) {
	c.L.Lock()
	done = true
	c.L.Unlock()

	// 唤醒任意一个等待的goroutine
	c.Signal()
}

func main() {
	m := &sync.Mutex{}
	c := sync.NewCond(m)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Soldiers(i, c, &wg)
	}

	go Waiter(c)

	Officer(c)
	wg.Wait()
	fmt.Println("所有大兵干完饭")

	CleanUp(c)
	time.Sleep(3 * time.Second)
	fmt.Println("打扫结束")
}
