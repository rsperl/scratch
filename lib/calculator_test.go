package lib

import (
	"math"
	"testing"
)

func TestCalculatorAdd(t *testing.T) {
	calc := NewCalculator()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed signs", 5, -3, 2},
		{"with zero", 5, 0, 5},
		{"decimals", 2.5, 3.7, 6.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculatorSubtract(t *testing.T) {
	calc := NewCalculator()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -5, -3, -2},
		{"mixed signs", 5, -3, 8},
		{"with zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculatorMultiply(t *testing.T) {
	calc := NewCalculator()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative numbers", -5, -3, 15},
		{"mixed signs", 5, -3, -15},
		{"with zero", 5, 0, 0},
		{"with one", 5, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculatorDivide(t *testing.T) {
	calc := NewCalculator()

	t.Run("valid division", func(t *testing.T) {
		result, err := calc.Divide(10, 2)
		if err != nil {
			t.Errorf("Divide(10, 2) returned unexpected error: %v", err)
		}
		if result != 5 {
			t.Errorf("Divide(10, 2) = %v; want 5", result)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calc.Divide(10, 0)
		if err == nil {
			t.Error("Divide(10, 0) should return an error")
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		result, err := calc.Divide(-10, 2)
		if err != nil {
			t.Errorf("Divide(-10, 2) returned unexpected error: %v", err)
		}
		if result != -5 {
			t.Errorf("Divide(-10, 2) = %v; want -5", result)
		}
	})
}

func TestCalculatorPower(t *testing.T) {
	calc := NewCalculator()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"2 to the power of 3", 2, 3, 8},
		{"5 to the power of 2", 5, 2, 25},
		{"any number to power 0", 5, 0, 1},
		{"any number to power 1", 5, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Power(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Power(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculatorSquareRoot(t *testing.T) {
	calc := NewCalculator()

	t.Run("valid square root", func(t *testing.T) {
		result, err := calc.SquareRoot(16)
		if err != nil {
			t.Errorf("SquareRoot(16) returned unexpected error: %v", err)
		}
		if result != 4 {
			t.Errorf("SquareRoot(16) = %v; want 4", result)
		}
	})

	t.Run("square root of zero", func(t *testing.T) {
		result, err := calc.SquareRoot(0)
		if err != nil {
			t.Errorf("SquareRoot(0) returned unexpected error: %v", err)
		}
		if result != 0 {
			t.Errorf("SquareRoot(0) = %v; want 0", result)
		}
	})

	t.Run("negative number", func(t *testing.T) {
		_, err := calc.SquareRoot(-4)
		if err == nil {
			t.Error("SquareRoot(-4) should return an error")
		}
	})
}
