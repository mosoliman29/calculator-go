package operation

import "fmt"

// Add returns the sum of two numbers.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between two numbers.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers.
func Divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	return a / b, nil
}

// Calculate performs the selected operation.
func Calculate(a, b float64, operator string) (float64, error) {

	switch operator {

	case "+":
		return Add(a, b), nil

	case "-":
		return Subtract(a, b), nil

	case "*":
		return Multiply(a, b), nil

	case "/":
		return Divide(a, b)

	default:
		return 0, fmt.Errorf("invalid operator")
	}
}
