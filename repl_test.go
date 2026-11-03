package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"only whitespace": {
			input:    "   \t\n  ",
			expected: []string{},
		},
		"single word": {
			input:    "Pikachu",
			expected: []string{"pikachu"},
		},
		"assignment example 1": {
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		"assignment example 2": {
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		"tabs and newlines between words": {
			input:    "hello\tworld\nfoo",
			expected: []string{"hello", "world", "foo"},
		},
		"multiple spaces between words": {
			input:    "a     b",
			expected: []string{"a", "b"},
		},
		"unicode": {
			input:    "ÉCLAIR Pokémon",
			expected: []string{"éclair", "pokémon"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("want %v, got %v", tc.expected, got)
			}
		})
	}
}
