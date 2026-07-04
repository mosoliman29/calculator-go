package main

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
func ReadOperator(reader *bufio.Reader) (string, error) {

	fmt.Println("\nSelect operation")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Print(">>> ")

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	switch input {

	case "1", "+":
		return "+", nil

	case "2", "-":
		return "-", nil

	case "3", "*":
		return "*", nil

	case "4", "/":
		return "/", nil

	default:
		return "", fmt.Errorf("please choose 1-4 or type + - * /")
	}
}
