package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func query() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		println("get website failed. err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(string(body))

	// resp2, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	// if err != nil {
	// 	println("get website failed. err:", err)
	// 	return
	// }
	// defer resp2.Body.Close()
	// fmt.Println(resp2.Body)

	resp3, err := http.PostForm("http://example.com/form",
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		println("get website failed. err:", err)
		return
	}
	defer resp3.Body.Close()

	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(string(body3))
}

func queryAndParams() {
	apiURL := "http://example.com/api"

	data := url.Values{}
	data.Set("name", "jack")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiURL)
	if err != nil {
		fmt.Println("parse url request url failed, err:", err)
		return
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("get url failed, err:", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get response failed. err:", err)
		return
	}
	fmt.Println(string(b))
}

// server
func servGetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func postQueryParams() {
	url := "http://example.com/post"

	contentType := "application/json"
	data := `{"name": "jack", "age":"18"}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post query failed. err", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response content fail. err:", err)
		return
	}
	fmt.Println(string(b))
}

func servPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request body failed. err:", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func serverHandler() {

	http.HandleFunc("/foo", fooHandler)
	// http.HandlerFunc("/bar", func(w http.ResponseWriter, r *http.request) {
	// 	fmt.FPrintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println("http server failed. err:", err)
		return
	}
}

func main() {
	// query()

	// queryAndParams()
	// postQueryParams()

	serverHandler()
}
