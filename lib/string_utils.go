package lib

import (
	"strings"
)

// StringUtils provides string manipulation functions
type StringUtils struct{}

// NewStringUtils creates a new StringUtils instance
func NewStringUtils() *StringUtils {
	return &StringUtils{}
}

// Reverse reverses a string
func (s *StringUtils) Reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome
func (s *StringUtils) IsPalindrome(input string) bool {
	cleaned := strings.ToLower(strings.ReplaceAll(input, " ", ""))
	return cleaned == s.Reverse(cleaned)
}

// CountVowels counts the number of vowels in a string
func (s *StringUtils) CountVowels(input string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range input {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

// ToTitleCase converts a string to title case
func (s *StringUtils) ToTitleCase(input string) string {
	return strings.Title(strings.ToLower(input))
}
