package main

import (
	"fmt"
)

type Individual struct {
	Name    string
	Age     int
	Email   string
	Country string
}

func main() {
	group := []Individual{
		{Name: "Alice", Age: 30, Email: "alice@example.com", Country: "USA"},
		{Name: "Bob", Age: 25, Email: "bob@example.com", Country: "USA"},
		{Name: "Charlie", Age: 35, Email: "charlie@example.com", Country: "USA"},
	}

	for idx := 0; idx < len(group); idx++ {
		member := group[idx]
		fmt.Printf("Name: %s, Age: %d, Email: %s, Country: %s\n", member.Name, member.Age, member.Email, member.Country)
	}
}
