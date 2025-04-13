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

func TestCommandHelp(t *testing.T) {
	cfg := &config{}
	err := commandHelp(cfg)
	if err != nil {
		t.Errorf("commandHelp returned an error: %v", err)
	}
	// No specific output to check, but ensure no errors occur.
}

func TestCommandExit(t *testing.T) {
	cfg := &config{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("commandExit did not call os.Exit")
		}
	}()
	_ = commandExit(cfg)
}
