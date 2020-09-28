package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func printer() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5 task
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	// print result
	for a := 1; a <= 5; a++ {
		<-results
	}
}

func main() {
	printer()
}
