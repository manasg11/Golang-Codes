package main

import (
	"fmt"
)

func main() {
	var upperLimit int

	fmt.Print("Please enter the upper limit: ")
	fmt.Scan(&upperLimit)

	fmt.Println("Odd numbers:")

	for number := 1; number <= upperLimit; number += 2 {
		fmt.Print(number)
		fmt.Println()
	}
}
