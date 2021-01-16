package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Fluent inteface

// Point struct
type Point struct {
	Longitude     string
	Latitude      string
	Distance      string
	ElevationGain string
	ElevationLoss string
}

// Reader struct
type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func parse(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}

func main() {
	f, _ := os.Open("a")
	p, err := parse(f)
	if err != nil {
		fmt.Print(p)
	}
}
