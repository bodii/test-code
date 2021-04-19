package main

import (
	"context"
	"fmt"
)

func example() {
	ctx, cancle := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cancle()
			break
		}
	}
	cancle()
}

func main() {
	example()
}
