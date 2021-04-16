package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Data struct {
	Number string `json:"number"`
	Title  string `json:"title"`
	Heat   string `json:"heat"`
}

func main() {
	res, err := http.Get("https://s.weibo.com/top/summary")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data []Data

	dom.Find(".td-01").Each(func(i int, selection *goquery.Selection) {
		data = append(data, Data{Number: selection.Text()})
	})

	dom.Find(".td-02>a").Each(func(i int, selection *goquery.Selection) {
		data[i].Title = selection.Text()
	})

	dom.Find(".td-02>span").Each(func(i int, selection *goquery.Selection) {
		data[i].Heat = selection.Text()
	})

	res_json, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(res_json))
}
