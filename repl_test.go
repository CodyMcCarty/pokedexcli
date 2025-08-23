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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		// add more cases here
	}

	for _, test := range cases {
		actual := cleanInput(test.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(test.expected) != len(actual) {
			t.Errorf(`---------------------------------
Test Failed:
Inputs: "%v"
Expecting: 	(length = %v)
%v
Actual: 	(length = %v)
%v
Fail
`, test.input, len(test.expected), sliceWithBullets(test.expected), len(actual), sliceWithBullets(actual))
		} else {
			fmt.Printf(`---------------------------------
Test Passed:
Inputs: "%v"
Expecting: 	(length = %v)
%v
Actual: 	(length = %v)
%v
Pass
`, test.input, len(test.expected), sliceWithBullets(test.expected), len(actual), sliceWithBullets(actual))
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := test.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if actualWord != expectedWord {
				t.Errorf(`---------------------------------
Test Failed:
Inputs: "%v"
Expecting: 	%v
Actual: 	%v
Fail
`, test.input, expectedWord, actualWord)
			} else {
				fmt.Printf(`---------------------------------
Test Passed:
Inputs: "%v"
Expecting: 	%v
Actual: 	%v
Pass
`, test.input, expectedWord, actualWord)
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
