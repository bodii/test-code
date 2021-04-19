package main

import "fmt"

func example(chExit chan bool) {
EXIT:
	for {
		select {
		case v, ok := <-chExit:
			if !ok {
				fmt.Println("close channel 2", v)
				break EXIT
			}
			fmt.Println("ch2 val = ", v)
		}
	}

	fmt.Println("exit")
}

func main() {
	chExit := make(chan bool)
	example(chExit)
}
