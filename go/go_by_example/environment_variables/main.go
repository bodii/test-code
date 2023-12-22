package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("FOO:", os.Getenv("FOO"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
