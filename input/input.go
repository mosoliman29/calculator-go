package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// ReadNumber reads a floating-point number from the user.
func ReadNumber(reader *bufio.Reader, prompt string) (float64, bool) {

	for {

		fmt.Print(prompt)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return 0, true
		}

		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "exit") {
			return 0, true
		}

		number, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}

		return number, false
	}
}

// ReadOperator reads the operator from the user.
func ReadOperator(reader *bufio.Reader) (string, error, bool) {

	fmt.Println("\nSelect operation")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Exit")
	fmt.Print(">>> ")

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err, false
	}

	input = strings.TrimSpace(input)

	if strings.EqualFold(input, "exit") || input == "5" {
		return "", nil, true
	}

	switch input {

	case "1", "+":
		return "+", nil, false

	case "2", "-":
		return "-", nil, false

	case "3", "*":
		return "*", nil, false

	case "4", "/":
		return "/", nil, false

	default:
		return "", fmt.Errorf("please choose 1-4 or type + - * /"), false
	}
}
