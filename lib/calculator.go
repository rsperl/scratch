package lib

import (
	"errors"
	"math"
)

// Calculator provides basic arithmetic operations
type Calculator struct{}

// NewCalculator creates a new Calculator instance
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add returns the sum of two numbers
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference of two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Power returns a raised to the power of b
func (c *Calculator) Power(a, b float64) float64 {
	return math.Pow(a, b)
}

// SquareRoot returns the square root of a number
func (c *Calculator) SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}
	return math.Sqrt(a), nil
}
