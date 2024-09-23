package main

import (
	"fmt"
)

func computeFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * computeFactorial(n-1)
}

func main() {
	var input int
	fmt.Print("Please enter a non-negative integer: ")
	fmt.Scan(&input)

	if input < 0 {
		fmt.Println("Error: Factorial is not defined for negative numbers.")
		return
	}

	factorialResult := computeFactorial(input)
	fmt.Printf("The factorial of %d is %d.\n", input, factorialResult)
}
