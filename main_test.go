package main

import (
	"reflect"
	"testing"
)

func TestWordFrequencyCount(t *testing.T) {
	input := "Hello, world! Hello Go. Go Go."
	expected := map[string]int{
		"hello": 2,
		"world": 1,
		"go":    3,
	}
	result := WordFrequencyCount(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal, Panama", true},
		{"racecar", true},
		{"Hello, world", false},
		{"No 'x' in Nixon", true},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %v, got %v", test.input, test.expected, result)
		}
	}
}
