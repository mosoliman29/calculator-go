package operation

import "testing"

func TestAdd_Positive(t *testing.T) {
	if Add(5, 3) != 8 {
		t.Error("Add(5, 3) should be 8")
	}
}

func TestAdd_Negative(t *testing.T) {
	if Add(-5, -3) != -8 {
		t.Error("Add(-5, -3) should be -8")
	}
}

func TestAdd_Zero(t *testing.T) {
	if Add(0, 7) != 7 {
		t.Error("Add(0, 7) should be 7")
	}
}

func TestAdd_Float(t *testing.T) {
	if Add(1.5, 2.5) != 4.0 {
		t.Error("Add(1.5, 2.5) should be 4.0")
	}
}

func TestSubtract_Positive(t *testing.T) {
	if Subtract(10, 4) != 6 {
		t.Error("Subtract(10, 4) should be 6")
	}
}

func TestSubtract_NegativeResult(t *testing.T) {
	if Subtract(3, 10) != -7 {
		t.Error("Subtract(3, 10) should be -7")
	}
}

func TestSubtract_Zero(t *testing.T) {
	if Subtract(7, 0) != 7 {
		t.Error("Subtract(7, 0) should be 7")
	}
}

func TestMultiply_Positive(t *testing.T) {
	if Multiply(5, 6) != 30 {
		t.Error("Multiply(5, 6) should be 30")
	}
}

func TestMultiply_Negative(t *testing.T) {
	if Multiply(-4, 3) != -12 {
		t.Error("Multiply(-4, 3) should be -12")
	}
}

func TestMultiply_Zero(t *testing.T) {
	if Multiply(5, 0) != 0 {
		t.Error("Multiply(5, 0) should be 0")
	}
}

func TestMultiply_Float(t *testing.T) {
	if Multiply(2.5, 4) != 10.0 {
		t.Error("Multiply(2.5, 4) should be 10.0")
	}
}

func TestDivide_Exact(t *testing.T) {
	result, err := Divide(20, 4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %g", result)
	}
}

func TestDivide_FloatResult(t *testing.T) {
	result, err := Divide(7, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 3.5 {
		t.Errorf("expected 3.5, got %g", result)
	}
}

func TestDivide_Negative(t *testing.T) {
	result, err := Divide(-10, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != -5 {
		t.Errorf("expected -5, got %g", result)
	}
}

func TestDivide_ByZero(t *testing.T) {
	_, err := Divide(5, 0)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

func TestDivide_ZeroDividend(t *testing.T) {
	result, err := Divide(0, 5)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 0 {
		t.Errorf("expected 0, got %g", result)
	}
}

func TestCalculate_Addition(t *testing.T) {
	result, err := Calculate(5, 3, "+")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 8 {
		t.Errorf("expected 8, got %g", result)
	}
}

func TestCalculate_Subtraction(t *testing.T) {
	result, err := Calculate(10, 4, "-")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 6 {
		t.Errorf("expected 6, got %g", result)
	}
}

func TestCalculate_Multiplication(t *testing.T) {
	result, err := Calculate(5, 6, "*")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 30 {
		t.Errorf("expected 30, got %g", result)
	}
}

func TestCalculate_Division(t *testing.T) {
	result, err := Calculate(20, 4, "/")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %g", result)
	}
}

func TestCalculate_DivisionByZero(t *testing.T) {
	_, err := Calculate(5, 0, "/")
	if err == nil {
		t.Error("expected division by zero error")
	}
}

func TestCalculate_InvalidOperator(t *testing.T) {
	_, err := Calculate(1, 2, "%")
	if err == nil {
		t.Error("expected error for invalid operator")
	}
}
