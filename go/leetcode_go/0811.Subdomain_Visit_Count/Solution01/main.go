package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	val := make(map[string]int)
	rest := make([]string, 0)

	for _, v := range cpdomains {
		strArr := strings.Fields(v)
		visits, _ := strconv.Atoi(strArr[0])
		domains := strings.Split(strArr[1], ".")
		len := len(domains)

		if len == 3 {
			val[domains[2]] += visits
			parent := strings.Join([]string{domains[len-2], domains[len-1]}, ".")
			val[parent] += visits
		} else if len == 2 {
			val[domains[1]] += visits
		}
		val[strArr[1]] += visits
	}

	for k, v := range val {
		rest = append(rest, fmt.Sprintf("%d %s", v, k))
	}

	return rest
}

func test01() {
	cpdomains := []string{"9001 discuss.leetcode.com"}
	reslut := []string{"9001 discuss.leetcode.com",
		"9001 leetcode.com", "9001 com"}
	fmt.Printf("test01 cpdomains:= %v\n", cpdomains)
	fmt.Printf("success result:  %v \n", reslut)
	fmt.Println("result:", subdomainVisits(cpdomains))
	fmt.Println()
}

func test02() {
	cpdomains := []string{"900 google.mail.com",
		"50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	reslut := []string{"901 mail.com", "50 yahoo.com",
		"900 google.mail.com", "5 wiki.org", "5 org",
		"1 intel.mail.com", "951 com"}
	fmt.Printf("test02 cpdomains:= %v\n", cpdomains)
	fmt.Printf("success result:  %v \n", reslut)
	fmt.Println("result:", subdomainVisits(cpdomains))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
