package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func withBuffer(iterations int) {
	for i := 0; i < iterations; i++ {
		buf := bufPool.Get().(*bytes.Buffer)
		buf.Reset() // Clear the buffer before using
		buf.WriteString("Hello, world!")
		_ = buf.String()
		bufPool.Put(buf)
	}
}

func withoutBuffer(iterations int) {
	for i := 0; i < iterations; i++ {
		buf := new(bytes.Buffer)
		buf.WriteString("Hello, world!")
		_ = buf.String()
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program [withbuffer|withoutbuffer] iterations")
		os.Exit(1)
	}

	allocType := os.Args[1]
	iterations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid iterations number: %v\n", err)
		os.Exit(1)
	}

	switch allocType {
	case "withbuffer":
		withBuffer(iterations)
	case "withoutbuffer":
		withoutBuffer(iterations)
	default:
		fmt.Println("Invalid allocation type. Use 'withbuffer' or 'withoutbuffer'")
		os.Exit(1)
	}
}
