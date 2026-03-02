package lib

import (
	"testing"
)

func TestStringUtilsReverse(t *testing.T) {
	su := NewStringUtils()
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple word", "hello", "olleh"},
		{"palindrome", "racecar", "racecar"},
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"with unicode", "🎉🎊", "🎊🎉"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := su.Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringUtilsIsPalindrome(t *testing.T) {
	su := NewStringUtils()
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"simple palindrome", "racecar", true},
		{"not palindrome", "hello", false},
		{"palindrome with spaces", "race car", true},
		{"mixed case palindrome", "RaceCar", true},
		{"single char", "a", true},
		{"empty string", "", true},
		{"two same chars", "aa", true},
		{"two different chars", "ab", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := su.IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringUtilsCountVowels(t *testing.T) {
	su := NewStringUtils()
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"simple word", "hello", 2},
		{"all vowels", "aeiou", 5},
		{"all consonants", "bcdfg", 0},
		{"mixed case", "HeLLo WoRLd", 3},
		{"empty string", "", 0},
		{"uppercase vowels", "AEIOU", 5},
		{"with numbers", "h3ll0 w0rld", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := su.CountVowels(tt.input)
			if result != tt.expected {
				t.Errorf("CountVowels(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringUtilsToTitleCase(t *testing.T) {
	su := NewStringUtils()
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase word", "hello", "Hello"},
		{"multiple words", "hello world", "Hello World"},
		{"already title case", "Hello World", "Hello World"},
		{"all caps", "HELLO WORLD", "Hello World"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := su.ToTitleCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToTitleCase(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
