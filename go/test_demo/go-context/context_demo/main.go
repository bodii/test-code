package main

import "context"

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n:=1

		go func()(
			for {
				select {
				case <-ctx.Done():
					return 
				case dst <- n:
					n++
				}
			}
		)()

		return dst
	}
}
