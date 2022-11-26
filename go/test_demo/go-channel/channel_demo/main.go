package main

import (
	"fmt"
	"time"
)

const DateTime = "2006-01-02 15:04:05"

func main() {
	ch := make(chan string)
	go test(ch)
	for {
		select {
		case value, ok := <-ch:
			fmt.Println("value is:", value, " ok is:", ok)
			if !ok {
				return
			}
		case <-time.After(2 * time.Second):
			fmt.Println("time out ...")
			return
		}
	}
}

func test(ch chan string) {
	defer func() {
		fmt.Println("i am test, now close ch chan")
		close(ch)
	}()

	for i := 1; i < 5; i++ {
		ch <- time.Now().Format(DateTime)
		time.Sleep(1 * time.Second)
	}
}
