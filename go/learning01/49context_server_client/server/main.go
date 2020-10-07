package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func serverIndexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}

func server() {
	http.HandleFunc("/", serverIndexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	server()
}
