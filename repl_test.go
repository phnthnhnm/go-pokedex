package main

import "testing"

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
			input:    "  foo bar  ",
			expected: []string{"foo", "bar"},
		},
		{
			input:    "  a  b  c  ",
			expected: []string{"a", "b", "c"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d, but got length %d", c.input, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input '%s', expected '%s', but got '%s'", c.input, expectedWord, word)
			}
		}
	}

}
