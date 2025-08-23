package main

import (
	"fmt"
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
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "  Hello  WORLD  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// test length
		if len(actual) != len(c.expected) {
			t.Errorf(` Lengths should be equal
Inputs: "%v"
Expecting: 	(length = %v)
%v
Actual: 	(length = %v)
%v
`, c.input, len(c.expected), sliceWithBullets(c.expected), len(actual), sliceWithBullets(actual))
			continue
		}

		// test words
		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]
			if actualWord != expectedWord {
				t.Errorf(` Failed
Inputs: "%v"
Expecting: 	%v
Actual: 	%v
Fail
`, c.input, expectedWord, actualWord)
			}
		}
	}
}

func sliceWithBullets[T any](slice []T) string {
	output := ""
	for i, item := range slice {
		form := "  - %v\n"
		if i == (len(slice) - 1) {
			form = "  - %v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}
