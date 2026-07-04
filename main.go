package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========================")
	fmt.Println("     Go Calculator")
	fmt.Println("========================")
	fmt.Println("Type 'exit' at any time to quit.")

	for {

		fmt.Println()

		num1, exit := ReadNumber(reader, "Enter first number: ")
		if exit {
			fmt.Println("Goodbye!")
			return
		}

		num2, exit := ReadNumber(reader, "Enter second number: ")
		if exit {
			fmt.Println("Goodbye!")
			return
		}

		operator, err := ReadOperator(reader)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := Calculate(num1, num2, operator)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("\nResult: %g %s %g = %g\n", num1, operator, num2, result)
	}
}
