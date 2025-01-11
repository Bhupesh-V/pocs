package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Define a substantial struct to make the memory differences more noticeable
type LargeStruct struct {
	ID        int
	Name      string
	Data      [100]int    // Large array to increase struct size
	Metadata  [20]string  // More fields to make it realistic
	Timestamp time.Time
	Status    bool
}

// Stack allocation - returns struct by value
func stackAlloc(iterations int) LargeStruct {
	result := LargeStruct{
		ID:       1,
		Name:     "test",
		Status:   true,
		Timestamp: time.Now(),
	}
	
	// Fill the arrays to make sure they're used
	for i := 0; i < iterations && i < 100; i++ {
		result.Data[i] = i
	}
	for i := 0; i < 20; i++ {
		result.Metadata[i] = fmt.Sprintf("meta-%d", i)
	}
	
	return result // Copies the entire struct
}

// Heap allocation - returns pointer to struct
func heapAlloc(iterations int) *LargeStruct {
	result := LargeStruct{
		ID:       1,
		Name:     "test",
		Status:   true,
		Timestamp: time.Now(),
	}
	
	// Fill the arrays to make sure they're used
	for i := 0; i < iterations && i < 100; i++ {
		result.Data[i] = i
	}
	for i := 0; i < 20; i++ {
		result.Metadata[i] = fmt.Sprintf("meta-%d", i)
	}
	
	return &result // Returns pointer, forcing heap allocation
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

	start := time.Now()
	
	switch allocType {
	case "stack":
		result := stackAlloc(iterations)
		fmt.Printf("Stack result - ID: %d, Name: %s, First Data: %d\n", 
			result.ID, result.Name, result.Data[0])
	case "heap":
		result := heapAlloc(iterations)
		fmt.Printf("Heap result - ID: %d, Name: %s, First Data: %d\n", 
			result.ID, result.Name, result.Data[0])
	default:
		fmt.Println("Invalid allocation type. Use 'stack' or 'heap'")
		os.Exit(1)
	}

	duration := time.Since(start)
	fmt.Printf("Duration: %v\n", duration)
}