package main

import (
	"fmt"
)

func main() {
	for {
		var num1, num2 int
		var operation int

		fmt.Print("Please enter the first number: ")
		fmt.Scan(&num1)

		fmt.Print("Please enter the second number: ")
		fmt.Scan(&num2)

		fmt.Println("Select an operation:")
		fmt.Println("1: Addition")
		fmt.Println("2: Subtraction")
		fmt.Println("3: Multiplication")
		fmt.Println("4: Division")
		fmt.Println("5: Exit")
		fmt.Print("Enter your choice (1-5): ")
		fmt.Scan(&operation)

		if operation == 5 {
			fmt.Println("Terminating the program...")
			break
		}

		var outcome float64

		switch operation {
		case 1:
			outcome = float64(num1 + num2)
		case 2:
			outcome = float64(num1 - num2)
		case 3:
			outcome = float64(num1 * num2)
		case 4:
			if num2 != 0 {
				outcome = float64(num1) / float64(num2)
			} else {
				fmt.Println("Error: Cannot divide by zero")
				continue
			}
		default:
			fmt.Println("Invalid option selected")
			continue
		}

		fmt.Printf("The result is: %.2f\n", outcome)
	}
}
