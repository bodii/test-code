package main

import (
	"fmt"
	"learningcache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	learningcache.NewGroup("scores", 2<<10, learningcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "localhost:9999"
	peers := learningcache.NewHTTPPool(addr)
	log.Println("learningcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
