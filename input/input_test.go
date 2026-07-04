package input

import (
	"bufio"
	"errors"
	"strings"
	"testing"
)

func TestReadNumber_ValidNumber(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("42\n"))
	num, exit := ReadNumber(reader, "")
	if exit {
		t.Error("expected exit=false")
	}
	if num != 42 {
		t.Errorf("expected 42, got %g", num)
	}
}

func TestReadNumber_NegativeNumber(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("-10\n"))
	num, exit := ReadNumber(reader, "")
	if exit {
		t.Error("expected exit=false")
	}
	if num != -10 {
		t.Errorf("expected -10, got %g", num)
	}
}

func TestReadNumber_Float(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("3.14\n"))
	num, exit := ReadNumber(reader, "")
	if exit {
		t.Error("expected exit=false")
	}
	if num != 3.14 {
		t.Errorf("expected 3.14, got %g", num)
	}
}

func TestReadNumber_Exit(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("exit\n"))
	_, exit := ReadNumber(reader, "")
	if !exit {
		t.Error("expected exit=true for 'exit'")
	}
}

func TestReadNumber_ExitCaseInsensitive(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("EXIT\n"))
	_, exit := ReadNumber(reader, "")
	if !exit {
		t.Error("expected exit=true for 'EXIT'")
	}
}

func TestReadNumber_InvalidThenValid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("abc\n42\n"))
	num, exit := ReadNumber(reader, "")
	if exit {
		t.Error("expected exit=false after valid input")
	}
	if num != 42 {
		t.Errorf("expected 42, got %g", num)
	}
}

func TestReadNumber_ReaderError(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	_, exit := ReadNumber(reader, "")
	if !exit {
		t.Error("expected exit=true on reader error")
	}
}

func TestReadOperator_AdditionByNumber(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("1\n"))
	op, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if op != "+" {
		t.Errorf("expected '+', got '%s'", op)
	}
}

func TestReadOperator_AdditionBySymbol(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("+\n"))
	op, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if op != "+" {
		t.Errorf("expected '+', got '%s'", op)
	}
}

func TestReadOperator_Subtraction(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("-\n"))
	op, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if op != "-" {
		t.Errorf("expected '-', got '%s'", op)
	}
}

func TestReadOperator_Multiplication(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("*\n"))
	op, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if op != "*" {
		t.Errorf("expected '*', got '%s'", op)
	}
}

func TestReadOperator_Division(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("/\n"))
	op, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if op != "/" {
		t.Errorf("expected '/', got '%s'", op)
	}
}

func TestReadOperator_AllChoices(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1\n", "+"},
		{"+\n", "+"},
		{"2\n", "-"},
		{"-\n", "-"},
		{"3\n", "*"},
		{"*\n", "*"},
		{"4\n", "/"},
		{"/\n", "/"},
	}
	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		op, err, exit := ReadOperator(reader)
		if exit {
			t.Errorf("ReadOperator(%q) unexpected exit", tt.input)
		}
		if err != nil {
			t.Errorf("ReadOperator(%q) unexpected error: %v", tt.input, err)
		}
		if op != tt.expected {
			t.Errorf("ReadOperator(%q) = %q, want %q", tt.input, op, tt.expected)
		}
	}
}

func TestReadOperator_ExitByKeyword(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("exit\n"))
	_, _, exit := ReadOperator(reader)
	if !exit {
		t.Error("expected exit=true for 'exit'")
	}
}

func TestReadOperator_ExitByKeywordCaseInsensitive(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("EXIT\n"))
	_, _, exit := ReadOperator(reader)
	if !exit {
		t.Error("expected exit=true for 'EXIT'")
	}
}

func TestReadOperator_ExitByNumber(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("5\n"))
	_, _, exit := ReadOperator(reader)
	if !exit {
		t.Error("expected exit=true for '5'")
	}
}

func TestReadOperator_Invalid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("x\n"))
	_, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err == nil {
		t.Error("expected error for invalid operator")
	}
}

func TestReadOperator_EmptyInput(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("\n"))
	_, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err == nil {
		t.Error("expected error for empty input")
	}
}

func TestReadOperator_ReaderError(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	_, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err == nil {
		t.Error("expected error on empty reader")
	}
}

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) {
	return 0, errors.New("mock error")
}

func TestReadNumber_ReaderWithError(t *testing.T) {
	reader := bufio.NewReader(errorReader{})
	_, exit := ReadNumber(reader, "")
	if !exit {
		t.Error("expected exit=true")
	}
}

func TestReadOperator_ReaderWithError(t *testing.T) {
	reader := bufio.NewReader(errorReader{})
	_, err, exit := ReadOperator(reader)
	if exit {
		t.Error("expected exit=false")
	}
	if err == nil {
		t.Error("expected error")
	}
}
