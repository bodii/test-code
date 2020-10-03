package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func demo1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k:=%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}

	wg.Wait()
}

/*
	go语言的sync包中提供了一个开箱即用的并发安全版map - sync.map
	开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。
	同时sync.Map 内置了诸如Store、Load、LoadOrStore、Delete、
	Range等操作方法。
*/
var m2 = sync.Map{}

func demo2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("k:=%v, v:=%v\n", key, value)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func main() {
	demo1()
	fmt.Println("------------------------------")
	demo2()
}
