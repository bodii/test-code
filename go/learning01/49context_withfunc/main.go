package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return ended goroutine
			case dst <- n:
				n++
			}

		}
	}()
	return dst
}

func testWithCancelDemo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func testWithDeadlineDemo() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

var wg sync.WaitGroup

func worker1(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("db connecting...")
		time.Sleep(time.Millisecond * 10) // takes 10 seconds
		select {
		case <-ctx.Done(): // auto execute after 50 milliseconds
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
}

func testWithTimeoutDemo() {
	// setting a timeout of 50 milliseconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()
	wg.Add(1)
	go worker1(ctx)
	time.Sleep(time.Second * 5)
	wg.Wait()
	fmt.Println("over")
}

// TraceCode type
type TraceCode string

func worker2(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	// on child goroutine get tracecode value
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}

	defer wg.Done()

LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}

	fmt.Println("worker done!")
}

func testWithValueDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512125522")

	wg.Add(1)
	go worker2(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}

func main() {
	// testWithCancelDemo()
	// testWithDeadlineDemo()
	// testWithTimeoutDemo()
	testWithValueDemo()

}
