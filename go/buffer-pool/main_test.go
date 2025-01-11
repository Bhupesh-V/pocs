package main

import (
	"testing"
)

func BenchmarkWithBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withBuffer(1000) // Adjust the number of iterations as needed
	}
}

func BenchmarkWithoutBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		withoutBuffer(1000) // Adjust the number of iterations as needed
	}
}
