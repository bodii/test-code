package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func example() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("status data type:%T, value:%v\n", result["status"], result["status"])
}

func main() {
	example()
}
