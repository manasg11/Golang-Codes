package main

import (
	"fmt"
	"strings"
)

func main() {
	var userAge int
	var isCitizen string

	fmt.Print("Enter your age: ")
	fmt.Scan(&userAge)

	fmt.Print("Are you a citizen of India? (yes/no): ")
	fmt.Scan(&isCitizen)

	isCitizen = strings.ToLower(isCitizen)

	if userAge >= 18 && isCitizen == "yes" {
		fmt.Println("Congratulations! You are eligible to vote.")
	} else {
		fmt.Println("Sorry, you are not eligible to vote.")
	}
}
