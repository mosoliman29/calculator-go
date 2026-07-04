package main

import "testing"

func TestAdd(t *testing.T) {

	if Add(5, 3) != 8 {
		t.Error("Add failed")
	}
}

func TestSubtract(t *testing.T) {

	if Subtract(10, 4) != 6 {
		t.Error("Subtract failed")
	}
}

func TestMultiply(t *testing.T) {

	if Multiply(5, 6) != 30 {
		t.Error("Multiply failed")
	}
}

func TestDivide(t *testing.T) {

	result, err := Divide(20, 4)

	if err != nil {
		t.Error(err)
	}

	if result != 5 {
		t.Error("Divide failed")
	}
}

func TestDivideByZero(t *testing.T) {

	_, err := Divide(5, 0)

	if err == nil {
		t.Error("Expected division by zero error")
	}
}
