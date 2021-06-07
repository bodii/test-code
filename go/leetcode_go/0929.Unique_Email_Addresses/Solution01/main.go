package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	address := make(map[string]int)
	for _, v := range emails {
		email := emailFormat(v)
		if _, ok := address[email]; !ok {
			address[email]++
		}
	}

	return len(address)
}

func emailFormat(email string) string {
	emailArr := strings.Split(email, "@")
	local := strings.Split(emailArr[0], "+")[0]
	subLocal := strings.Split(local, ".")
	localNew := strings.Join(subLocal, "")

	return localNew + "@" + emailArr[1]
}

func test01() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	success := 2

	fmt.Println("test01  emails:=", emails)
	fmt.Println("success result:", success)
	fmt.Println("result:", numUniqueEmails(emails))
	fmt.Println()
}

func test02() {
	emails := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	success := 3

	fmt.Println("test01  emails:=", emails)
	fmt.Println("success result:", success)
	fmt.Println("result:", numUniqueEmails(emails))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
