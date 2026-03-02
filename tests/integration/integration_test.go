package tests

import (
	"fmt"
	"testing"

	"github.com/rsperl/scratch/lib"
)

func TestHelloIntegration(t *testing.T) {
	expected := "hello integration"
	result := lib.HelloIntegration()
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

// TestCalculatorAndStringUtilsIntegration tests the interaction between Calculator and StringUtils
func TestCalculatorAndStringUtilsIntegration(t *testing.T) {
	calc := lib.NewCalculator()
	strUtil := lib.NewStringUtils()

	// Perform a calculation and convert result to string
	result := calc.Add(10, 20)
	resultStr := fmt.Sprintf("%.0f", result)

	// Check if the string representation is a palindrome
	// Note: "30" is not a palindrome, so we expect false
	isPalindrome := strUtil.IsPalindrome(resultStr)
	if isPalindrome {
		t.Errorf("Expected '30' not to be a palindrome")
	}

	// Test with a calculation that results in a palindrome-like number
	result2 := calc.Add(11, 0)
	resultStr2 := fmt.Sprintf("%.0f", result2)
	isPalindrome2 := strUtil.IsPalindrome(resultStr2)
	if !isPalindrome2 {
		t.Errorf("Expected '11' to be a palindrome")
	}
}

// TestCalculatorChainOperations tests chaining multiple calculator operations
func TestCalculatorChainOperations(t *testing.T) {
	calc := lib.NewCalculator()

	// Perform a series of operations: (10 + 5) * 2 - 3 = 27
	step1 := calc.Add(10, 5)        // 15
	step2 := calc.Multiply(step1, 2) // 30
	result := calc.Subtract(step2, 3) // 27

	expected := 27.0
	if result != expected {
		t.Errorf("Chain operations result = %v; want %v", result, expected)
	}
}

// TestStringUtilsWorkflow tests a complete workflow with string operations
func TestStringUtilsWorkflow(t *testing.T) {
	strUtil := lib.NewStringUtils()

	input := "hello world"

	// Reverse the string
	reversed := strUtil.Reverse(input)
	expectedReversed := "dlrow olleh"
	if reversed != expectedReversed {
		t.Errorf("Reverse operation failed: got %q; want %q", reversed, expectedReversed)
	}

	// Count vowels in original
	vowelCount := strUtil.CountVowels(input)
	expectedVowels := 3
	if vowelCount != expectedVowels {
		t.Errorf("CountVowels failed: got %d; want %d", vowelCount, expectedVowels)
	}

	// Convert to title case
	titleCase := strUtil.ToTitleCase(input)
	expectedTitle := "Hello World"
	if titleCase != expectedTitle {
		t.Errorf("ToTitleCase failed: got %q; want %q", titleCase, expectedTitle)
	}
}

// TestCalculatorComplexOperations tests more complex calculator operations
func TestCalculatorComplexOperations(t *testing.T) {
	calc := lib.NewCalculator()

	// Test: sqrt(16) + power(2, 3) = 4 + 8 = 12
	sqrt, err := calc.SquareRoot(16)
	if err != nil {
		t.Fatalf("SquareRoot failed: %v", err)
	}

	power := calc.Power(2, 3)
	result := calc.Add(sqrt, power)

	expected := 12.0
	if result != expected {
		t.Errorf("Complex operations result = %v; want %v", result, expected)
	}
}

// TestErrorHandlingIntegration tests error handling across operations
func TestErrorHandlingIntegration(t *testing.T) {
	calc := lib.NewCalculator()

	// Test division by zero error
	_, err := calc.Divide(10, 0)
	if err == nil {
		t.Error("Division by zero should return an error")
	}

	// Test negative square root error
	_, err = calc.SquareRoot(-1)
	if err == nil {
		t.Error("Square root of negative number should return an error")
	}

	// Verify valid operations still work
	result := calc.Add(5, 5)
	if result != 10 {
		t.Errorf("Valid operation after errors failed: got %v; want 10", result)
	}
}
