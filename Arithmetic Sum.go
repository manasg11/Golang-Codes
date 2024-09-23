package main

import (
	"fmt"
)

func main() {
	var upperLimit int
	total := 0

	fmt.Print("Please enter the upper limit: ")
	fmt.Scan(&upperLimit)

	for number := 1; number <= upperLimit; number++ {
		total += number
	}

	fmt.Printf("The sum of all numbers from 1 to %d is: %d\n", upperLimit, total)
}
