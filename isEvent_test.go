package main

import "testing"

func TestIsEven(t *testing.T) {

	testTable := []struct {
		numbers  int
		expected bool
	}{
		{
			numbers:  2,
			expected: true,
		},
		{
			numbers:  3,
			expected: false,
		},
	}

	for _, testCase := range testTable {

		result := IsEven(testCase.numbers)

		if result == testCase.expected {
			t.Error("Not even")
		}

	}
}
