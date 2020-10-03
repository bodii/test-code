package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("args is empty")
		return
	}

	for index, arg := range os.Args {
		fmt.Printf("args[%d]=%v\n", index, arg)
	}
}
