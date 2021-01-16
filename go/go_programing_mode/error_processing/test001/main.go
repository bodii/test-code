package main

import (
	"io"
	"log"
	"os"
)

// Close func
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// Open open file func
func Open(f string) (r io.Closer, err error) {
	r, err = os.Open(f)
	if err != nil {
		log.Fatal("open file fail")
	}

	return
}

func test01() {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}

	defer Close(r)

	r, err = Open("b")
	if err != nil {
		log.Fatalf("error opening 'b'\n")
	}
	defer Close(r)
}

func main() {
	test01()
}
