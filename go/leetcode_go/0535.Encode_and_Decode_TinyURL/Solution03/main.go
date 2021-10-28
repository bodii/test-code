package main

import (
	"fmt"
	"strconv"
)

const (
	hostDomain = "http://tinyurl.com/"
	baseChars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Codec struct {
	urls  map[int]string
	count int
}

func Constructor() Codec {
	return Codec{urls: make(map[int]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.count++
	this.urls[this.count] = longUrl
	return hostDomain + strconv.Itoa(this.count)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	key, _ := strconv.Atoi(shortUrl[len(hostDomain):])
	return this.urls[key]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

func test01() {
	logUrl := "https://leetcode.com/problems/design-tinyurl"
	obj := Constructor()
	url := obj.encode(logUrl)
	ans := obj.decode(url)
	successResult := "http://tinyurl.com/4e9iAk"

	fmt.Println("test01 logUrl:=", logUrl)
	fmt.Println("success encode result:=", successResult)
	fmt.Println("encodeed url result:=", url)
	fmt.Println("decode result:=", ans)
	fmt.Println()

}

func main() {
	test01()
}
