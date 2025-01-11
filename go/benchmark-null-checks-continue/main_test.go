package main

import (
	"testing"
)

// Generate a large test slice with a mix of nil and non-nil values
func generateLargeTestSlice(size int) []*int {
	arr := make([]*int, size)
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			val := i
			arr[i] = &val
		} else {
			arr[i] = nil
		}
	}
	return arr
}

func BenchmarkSumWithContinue(b *testing.B) {
	arr := generateLargeTestSlice(10_000_000) // Large dataset
	for i := 0; i < b.N; i++ {
		SumWithContinue(arr)
	}
}

func BenchmarkSumWithIfCheck(b *testing.B) {
	arr := generateLargeTestSlice(10_000_000) // Large dataset
	for i := 0; i < b.N; i++ {
		SumWithIfCheck(arr)
	}
}
