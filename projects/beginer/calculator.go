package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	var operation int

	// Function definitions
	add := func(a, b int) int {
		return a + b
	}

	subtract := func(a, b int) int {
		return a - b
	}

	multiply := func(a, b int) int {
		return a * b
	}

	divide := func(a, b int) int {
		return a / b
	}

	fmt.Println("Enter first number:")
	_, err := fmt.Scanln(&number1)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Enter second number:")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Choose operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var result int
	switch operation {
	case 1:
		result = add(number1, number2)
		fmt.Printf("Result: %d\n", result)
	case 2:
		result = subtract(number1, number2)
		fmt.Printf("Result: %d\n", result)
	case 3:
		result = multiply(number1, number2)
		fmt.Printf("Result: %d\n", result)
	case 4:
		if number2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		result = divide(number1, number2)
		fmt.Printf("Result: %d\n", result)
	default:
		fmt.Println("Invalid operation")
	}
}
