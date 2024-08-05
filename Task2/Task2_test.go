package main

import (
	"reflect"
	"testing"
)

func TestfreqCount(t *testing.T) {
	text := "Artificial intelligence and machine learning are transforming the future of technology, as artificial intelligence continues to advance."
	expected := map[string]int{
		"artificial":     2,
		"intelligence":   2,
		"and":            1,
		"machine":        1,
		"learning":       1,
		"are":            1,
		"transforming":   1,
		"the":            1,
		"future":         1,
		"of":             1,
		"technology":     1,
		"as":             1,
		"continues":      1,
		"to":             1,
		"advance":        1,
	}


	result := freqCount(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"No lemon, no melon", true},
		{"racecar", true},
		{"hello", false},
		{"", true},
		{"Was it a car or a cat I saw", true},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("For input \"%s\", expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
