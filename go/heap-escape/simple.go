package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Stack allocation - returns value
func stackAlloc(iterations int) int {
	var result int
	for i := 0; i < iterations; i++ {
		x := 42
		result += x
	}
	return result
}

// Heap allocation - returns pointer
func heapAlloc(iterations int) *int {
	var result int
	for i := 0; i < iterations; i++ {
		x := 42
		result += x
	}
	return &result // Forces heap allocation
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program [stack|heap] iterations")
		os.Exit(1)
	}

	allocType := os.Args[1]
	iterations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid iterations number: %v\n", err)
		os.Exit(1)
	}

	// Run the specified allocation type
	start := time.Now()

	switch allocType {
	case "stack":
		result := stackAlloc(iterations)
		fmt.Printf("Stack result: %d\n", result)
	case "heap":
		result := heapAlloc(iterations)
		fmt.Printf("Heap result: %d\n", *result)
	default:
		fmt.Println("Invalid allocation type. Use 'stack' or 'heap'")
		os.Exit(1)
	}

	duration := time.Since(start)
	fmt.Printf("Duration: %v\n", duration)
}
