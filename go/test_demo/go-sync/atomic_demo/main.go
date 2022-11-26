package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func example1() {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))

	var mu sync.Mutex
	read := func(key string) (value string) {
		m1 := m.Load().(Map)
		return m1[key]
	}

	insert := func(key, val string) {
		mu.Lock()
		defer mu.Unlock()

		m1 := m.Load().(Map)
		m2 := make(Map)
		for k, v := range m1 {
			m2[k] = v
		}

		m2[key] = val
		m.Store(m2)
	}

	_, _ = read, insert
}

func example2() {
	config := atomic.Value{}
	config.Store(22)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			if i == 0 {
				config.Store(23)
			}

			fmt.Println(config.Load())
		}(i)
	}
	wg.Wait()
}

func example3() {
	var opts int64 = 0
	for i := 0; i < 50; i++ {
		atomic.AddInt64(&opts, 3)
	}
	time.Sleep(1 * time.Second)

	fmt.Println("opts: ", atomic.LoadInt64(&opts))
}

func example4() {
	var a, b int32 = 13, 13
	var c int32 = 9

	res := atomic.CompareAndSwapInt32(&a, b, c)
	fmt.Println("swapped: ", res)
	fmt.Println("replaceing : ", c)
	fmt.Println("replaced :", a)
}

func main() {
	// example1()
	// example2()
	// example3()
	example4()
}
