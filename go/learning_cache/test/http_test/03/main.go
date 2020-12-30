package main

import (
	"flag"
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

func createGroup() *learningcache.Group {
	return learningcache.NewGroup("scores", 2<<10, learningcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, lg *learningcache.Group) {
	peers := learningcache.NewHTTPPool(addr)
	peers.Set(addrs...)
	lg.RegisterPeers(peers)
	log.Println("learningcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, lg *learningcache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			key := req.URL.Query().Get("key")
			view, err := lg.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		}))
}

func main() {
	var port int
	var api bool

	flag.IntVar(&port, "port", 8001, "learningcache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	l := createGroup()
	if api {
		go startAPIServer(apiAddr, l)
	}
	startCacheServer(addrMap[port], []string(addrs), l)
}
