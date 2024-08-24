package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func welcome() {
	fmt.Println("Starting Calculator!")
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func parseInput(input string) (float64, float64, string, error) {
	tokens := strings.Fields(input)
	if len(tokens) != 3 {
		return 0, 0, "", fmt.Errorf("invalid input format")
	}

	num1, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid number: %s", tokens[0])
	}

	num2, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid number: %s", tokens[2])
	}

	operator := tokens[1]
	return num1, num2, operator, nil
}

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return add(num1, num2), nil
	case "-":
		return subtract(num1, num2), nil
	case "*":
		return multiply(num1, num2), nil
	case "/":
		return divide(num1, num2)
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}

func main() {
	welcome()

	for {
		fmt.Print("Enter calculation (e.g., 3 + 4): ")
		var input string
		fmt.Scanln(&input)

		num1, num2, operator, err := parseInput(input)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		result, err := calculate(num1, num2, operator)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}
}
