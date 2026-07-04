package app

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func runApp(t *testing.T, input string) string {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	oldStdout := os.Stdout
	rOut, wOut, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = wOut
	defer func() { os.Stdout = oldStdout }()

	Run()

	wOut.Close()
	var buf bytes.Buffer
	io.Copy(&buf, rOut)
	return buf.String()
}

func TestRun_ExitImmediately(t *testing.T) {
	output := runApp(t, "exit\n")
	if !strings.Contains(output, "Goodbye!") {
		t.Errorf("expected 'Goodbye!' in output, got:\n%s", output)
	}
}

func TestRun_ExitOnSecondNumber(t *testing.T) {
	output := runApp(t, "42\nexit\n")
	if !strings.Contains(output, "Goodbye!") {
		t.Errorf("expected 'Goodbye!' in output, got:\n%s", output)
	}
}

func TestRun_AllOperations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Addition", "5\n3\n+\nexit\n", "Result: 5 + 3 = 8"},
		{"Subtraction", "10\n4\n-\nexit\n", "Result: 10 - 4 = 6"},
		{"Multiplication", "5\n6\n*\nexit\n", "Result: 5 * 6 = 30"},
		{"Division", "20\n4\n/\nexit\n", "Result: 20 / 4 = 5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := runApp(t, tt.input)
			if !strings.Contains(output, tt.expected) {
				t.Errorf("expected %q in output, got:\n%s", tt.expected, output)
			}
		})
	}
}

func TestRun_InvalidOperator(t *testing.T) {
	output := runApp(t, "5\n3\nx\nexit\n")
	if !strings.Contains(output, "please choose 1-4") {
		t.Errorf("expected operator error message in output, got:\n%s", output)
	}
	if !strings.Contains(output, "Goodbye!") {
		t.Errorf("expected 'Goodbye!' after recovering, got:\n%s", output)
	}
}

func TestRun_DivisionByZero(t *testing.T) {
	output := runApp(t, "5\n0\n/\nexit\n")
	if !strings.Contains(output, "division by zero") {
		t.Errorf("expected division by zero error in output, got:\n%s", output)
	}
	if !strings.Contains(output, "Goodbye!") {
		t.Errorf("expected 'Goodbye!' after recovering, got:\n%s", output)
	}
}

func TestRun_MultipleCalculationsThenExit(t *testing.T) {
	output := runApp(t, "10\n5\n+\n3\n4\n*\nexit\n")
	if !strings.Contains(output, "Result: 10 + 5 = 15") {
		t.Errorf("expected first result, got:\n%s", output)
	}
	if !strings.Contains(output, "Result: 3 * 4 = 12") {
		t.Errorf("expected second result, got:\n%s", output)
	}
}

func TestRun_InvalidOperatorThenValidOperator(t *testing.T) {
	output := runApp(t, "5\n3\nx\n10\n2\n+\nexit\n")
	if !strings.Contains(output, "Result: 10 + 2 = 12") {
		t.Errorf("expected result after recovery, got:\n%s", output)
	}
}

func TestRun_DivisionByZeroThenValidDivision(t *testing.T) {
	output := runApp(t, "5\n0\n/\n10\n2\n/\nexit\n")
	if !strings.Contains(output, "Result: 10 / 2 = 5") {
		t.Errorf("expected result after recovery, got:\n%s", output)
	}
}

func TestRun_FloatNumbers(t *testing.T) {
	output := runApp(t, "1.5\n2.5\n+\nexit\n")
	if !strings.Contains(output, "Result: 1.5 + 2.5 = 4") {
		t.Errorf("expected float result, got:\n%s", output)
	}
}

func TestRun_NegativeNumbers(t *testing.T) {
	output := runApp(t, "-10\n-5\n+\nexit\n")
	if !strings.Contains(output, "Result: -10 + -5 = -15") {
		t.Errorf("expected negative result, got:\n%s", output)
	}
}

func TestRun_ExitFromOperator(t *testing.T) {
	output := runApp(t, "5\n3\nexit\n")
	if !strings.Contains(output, "Goodbye!") {
		t.Errorf("expected 'Goodbye!' in output, got:\n%s", output)
	}
}

func TestRun_ExitFromOperatorByNumber(t *testing.T) {
	output := runApp(t, "5\n3\n5\n")
	if !strings.Contains(output, "Goodbye!") {
		t.Errorf("expected 'Goodbye!' in output, got:\n%s", output)
	}
}

func TestRun_ShowsBanner(t *testing.T) {
	output := runApp(t, "exit\n")
	if !strings.Contains(output, "Go Calculator") {
		t.Errorf("expected banner in output, got:\n%s", output)
	}
}
