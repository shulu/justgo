package gotest

import (
	"testing"
)

func TestFib(t *testing.T) {
	var fibTests = []struct {
		in       int //input
		expected int //expected result
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{7, 13},
	}
	for _, tt := range fibTests {
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("Fib(%d) = %d: expected %d", tt.in, actual, tt.expected)
		}
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}
