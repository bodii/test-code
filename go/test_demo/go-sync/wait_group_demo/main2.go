package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("Depositing %d to account with blance %d\n", value, balance)
	balance += value
}

func withDraw(value int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("Withdrawing %d from account with blance %d\n", value, balance)
	balance -= value
}

func main() {
	fmt.Println("Hello World")

	balance = 1000
	var wg sync.WaitGroup
	wg.Add(2)
	go withDraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}
