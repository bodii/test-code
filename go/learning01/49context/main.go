package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}

	wg.Done()
}

func testDemo1() {
	wg.Add(1)
	go worker()
	time.Sleep(time.Second * 3)
	exit = true
	wg.Wait()
	fmt.Println("over")
}

func worker2(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
	wg.Done()
}

func testDemo2() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker2(exitChan)
	time.Sleep(time.Second * 3)
	exitChan <- struct{}{}
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}

func worker3(ctx context.Context) {
	defer wg.Done()

	go worker3Chlid(ctx)

LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func worker3Chlid(ctx context.Context) {
LOOP:
	for {
		fmt.Println("child worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func testDemo3() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker3(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

func main() {
	// testDemo1()
	// testDemo2()
	testDemo3()
}
