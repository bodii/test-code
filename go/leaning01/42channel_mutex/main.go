package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则
	在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时
	等待一个锁时，唤醒的策略是随机的。
*/

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func mutexDemo() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

// 读写锁
/*
读写锁在Go语言中使用sync包中的RWMutex类型。
读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁
会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是
获取读锁还是写锁都会等待。
*/

var (
	y      int64
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock() // 互斥锁
	rwlock.Lock() // 加写锁
	y = y + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	wg.Done()
}

func read() {
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	wg.Done()
}

func readWriteLockDemo() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))
}

func main() {
	mutexDemo()

	readWriteLockDemo()
}
