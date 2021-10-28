package main

import (
	"fmt"
	"strings"
)

const (
	hostDomain = "http://tinyurl.com/"
	baseChars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Codec struct {
	urls  map[string]string
	count int
}

func Constructor() Codec {
	return Codec{urls: make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.count++
	key := this.getKey()
	this.urls[key] = longUrl
	return hostDomain + key
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.urls[shortUrl[len(hostDomain):]]
}

func (this *Codec) getKey() string {
	key := strings.Builder{}
	len := len(baseChars)

	for count := this.count; count > 0; count-- {
		key.WriteByte(baseChars[count%len])
	}

	return key.String()
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
