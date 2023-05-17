package fib_test

import (
	"testing"
	"github.com/andyantrim/perf/fib"
)

func TestFibRecursive(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},
	}

	for _, tt := range tests {
		got := fib.FibRecursive(tt.n)
		if got != tt.want {
			t.Errorf("FibRecursive(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

// And the same for iterative
func TestFibIterative(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},
	}

	for _, tt := range tests {
		got := fib.FibIterative(tt.n)
		if got != tt.want {
			t.Errorf("FibRecursive(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func BenchmarkFibRecursive(b *testing.B) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},
	}
	
	b.ResetTimer()
	for _, tt := range tests {
		got := fib.FibRecursive(tt.n)
		if got != tt.want {
			b.Errorf("FibRecursive(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

// And the same for iterative
func BenchmarkFibIterative(b *testing.B) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},
	}
	
	b.ResetTimer()
	for _, tt := range tests {
		got := fib.FibIterative(tt.n)
		if got != tt.want {
			b.Errorf("FibRecursive(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}
