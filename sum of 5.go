package main

import (
	"fmt"
)

func main() {
	var subject1, subject2, subject3, subject4, subject5 float64

	fmt.Println("Please enter the marks for five subjects:")

	fmt.Print("Marks for Subject 1: ")
	fmt.Scan(&subject1)

	fmt.Print("Marks for Subject 2: ")
	fmt.Scan(&subject2)

	fmt.Print("Marks for Subject 3: ")
	fmt.Scan(&subject3)

	fmt.Print("Marks for Subject 4: ")
	fmt.Scan(&subject4)

	fmt.Print("Marks for Subject 5: ")
	fmt.Scan(&subject5)

	total := subject1 + subject2 + subject3 + subject4 + subject5

	fmt.Printf("Total marks obtained: %.2f\n", total)
}
