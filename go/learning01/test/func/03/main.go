package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string {
		"Matthew", "Sarah", "Augustus", "Heidi",
		"Emilie", "Peter", "Giana", "Adriano", 
		"Aaron", "Elizabeth",
	}
	
	distribution = make(map[string]int, len(users))
)


func main() {
	left := dispatchCoin()
	fmt.Println("剩下: ", left)
}

func dispatchCoin() (left int) {
	count := 0
	for _, user := range users {
		distribution[user] = 0
		userL := strings.ToLower(user)
		userInENum := strings.Count(userL, "e")
		userInINum := strings.Count(userL, "i") * 2
		userInONum := strings.Count(userL, "o") * 3
		userInUNum := strings.Count(userL, "u") * 4
		userCount := userInENum + userInINum + userInONum + userInUNum
		count += userCount
		distribution[user] = userCount
	}
	fmt.Println(distribution)
	return coins - count
}

