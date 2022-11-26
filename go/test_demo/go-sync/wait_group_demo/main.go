package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)
	fmt.Println("Finished Executing Goroutine")
}

func example1() {
	fmt.Println("Go WaitGroup Tutorial")

	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()

	fmt.Println("Finished Executing my go program")
}

var urls = []string{
	"https://cn.bing.com",
	"https://golang.google.cn",
	"https://flutter.cn",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp.Status)

			wg.Done()
		}(url)
	}

	wg.Wait()
}

func example2() {
	fmt.Println("Go WaitGroup Tutorial")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// example1()
	example2()
}
