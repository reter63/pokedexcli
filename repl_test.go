package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "  heLlo  WoRld  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "Banana is a  Terrible  1st   breakfast!",
		expected: []string{"banana", "is", "a", "terrible", "1st", "breakfast!"},
	},
	{
		input:    "",
		expected: []string{},
	},
	{
		input:    "    ",
		expected: []string{},
	},
	{
		input:    "  -3  ",
		expected: []string{"-3"},
	},
}

for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice
	// if they don't match, use t.Errorf and continue to the next case
	if len(actual) != len(c.expected) {
		// error and continue here
		t.Errorf("Length does not match: \nExpected: %v\nActual: %v\n", len(c.expected), len(actual))
		continue
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if word != expectedWord {
			t.Errorf("Word does not match: \nExpected: %v\nActual: %v\n", expectedWord, word)
		}
	}
}
}
