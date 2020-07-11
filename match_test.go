package gofuzzy

import (
	"testing"
)

func TestFuzzyEquals(t *testing.T) {
	strA := "hello"
	strB := "helloo"
	result := FuzzyEquals(&strA, &strB, 0.0)
	if !result {
		t.Errorf("Expected true for hello v helloo. Result %t", result)
	}
}

func TestFuzzyMatch(t *testing.T) {
	strA := "hello"
	strB := "helloo"
	result := FuzzyMatch(&strA, &strB)
	if result < 0.001 {
		t.Errorf("Result = %f", result)
	}
}
