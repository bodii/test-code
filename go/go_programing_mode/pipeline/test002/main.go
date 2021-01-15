package main

import "fmt"

// EchoFunc func type
type EchoFunc func([]int) <-chan int

// PipeFunc func type
type PipeFunc func(<-chan int) <-chan int

func echo(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func gen(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			out <- i
		}
		close(out)
	}()

	return out
}

func odd(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			if i%2 != 0 {
				out <- i
			}
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		sum := 0
		for i := range in {
			sum += i
		}
		out <- sum
		close(out)
	}()

	return out
}

func pipeline(nums []int, echo EchoFunc, pipeFuns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for i := range pipeFuns {
		ch = pipeFuns[i](ch)
	}
	return ch
}

func test01() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range pipeline(nums, echo, odd, sq, sum) {
		fmt.Println(n)
	}
}

func main() {
	test01()
}
