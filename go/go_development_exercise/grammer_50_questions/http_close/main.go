package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func http_query_close_body() {
	req, err := http.Get("https://www.example.com")

	if req != nil {
		req.Body.Close()
	}

	checkError(err)

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	checkError(err)

	fmt.Println(string(body))
}

func http_query_close() {
	req, err := http.NewRequest("GET", "https://www.example.com", nil)
	checkError(err)

	req.Close = true
	// req.Header.Add("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		resp.Body.Close()
	}
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}

func http_transport() {
	transport := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &transport}

	resp, err := client.Get("https://www.example.com")
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)

	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(len(string(body)))
}

func main() {
	http_query_close_body()
	http_query_close()
	http_transport()
}
