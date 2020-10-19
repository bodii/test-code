package main

import "fmt"

func convertToTitle(n int) string {
	result := make([]byte, 0)
	for n > 0 {
		n--
		result = append([]byte{byte('A' + (n % 26))}, result...)
		n /= 26
	}
	return string(result)
}

func main() {
	fmt.Println(convertToTitle(701))
}
