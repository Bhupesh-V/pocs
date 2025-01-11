package main

import "fmt"

// SumWithContinue sums non-nil values using the "continue" approach
func SumWithContinue(arr []*int) int {
	sum := 0
	for _, val := range arr {
		if val == nil {
			continue
		}
		sum += *val
	}
	return sum
}

// SumWithIfCheck sums non-nil values using "if val != nil"
func SumWithIfCheck(arr []*int) int {
	sum := 0
	for _, val := range arr {
		if val != nil {
			sum += *val
		}
	}
	return sum
}

func main() {
	// Example usage with a smaller dataset
	a, b, c := 1, 2, 3
	arr := []*int{&a, nil, &b, nil, &c}
	fmt.Println("Sum with continue:", SumWithContinue(arr))
	fmt.Println("Sum with if check:", SumWithIfCheck(arr))
}
