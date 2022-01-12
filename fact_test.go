package main

import (
	"testing"
)

// Test function for iterative version of factorial function
func TestFactorial(t *testing.T) {
	var testcase = []struct {
		input, output int
	}{
		{6, 720},
		{0, 1},
		{1, 1},
		{2, 2},
	}
	for i, tt := range testcase {
		expectedOutput := factorial(tt.input)
		if expectedOutput != tt.output {
			t.Errorf("TEST[%d], failed.\nExpected: %v\nGot: %v\n", i, tt.output, expectedOutput)
		}
	}
}

// Test function for recursive version of factorial function
func TestFactorialRecursive(t *testing.T) {
	var testcase = []struct {
		input, output int
	}{
		{6, 720},
		{0, 1},
		{1, 1},
		{2, 2},
	}
	for i, tt := range testcase {
		expectedOutput := factRecursive(tt.input)
		if expectedOutput != tt.output {
			t.Errorf("TEST[%d], failed.\nExpected: %v\nGot: %v\n", i, tt.output, expectedOutput)
		}
	}
}

// Bench Marking Test function for iterative factorial function
func BenchmarkCheckFact(b *testing.B) {
	num := 10
	for i := 0; i < b.N; i++ {
		factorial(num)
	}
}

// Bench Marking Test function for recursive factorial function
func BenchmarkFactRecursive(b *testing.B) {
	num := 10
	for i := 0; i < b.N; i++ {
		factRecursive(num)
	}
}
