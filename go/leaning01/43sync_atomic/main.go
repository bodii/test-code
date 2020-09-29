package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对
	基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供
	的方法它在用户态就可以完成，因此性能比加锁操作更好。go语言中原子操作由内置的
	标准库go/atomic提供
*/
type counter interface {
	Inc()
	Load() int64
}

// CommonCounter 普通版
type CommonCounter struct {
	counter int64
}

// Inc increment
func (c CommonCounter) Inc() {
	c.counter++
}

// Load get
func (c CommonCounter) Load() int64 {
	return c.counter
}

// MutexCounter mutex
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

// Inc muterx inc
func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

// Load muterx load
func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// AtomicCounter struct
type AtomicCounter struct {
	counter int64
}

// Inc atomic inc
func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

// Load atomic load
func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	c1 := CommonCounter{} // not safe
	test(c1)

	c2 := MutexCounter{} // use mutex lock
	test(&c2)

	c3 := AtomicCounter{} // use atomic
	test(&c3)
}
