package main

import (
	"fmt"
	"sort"
)

func getMinMax(values []int) (int, int) {
	minValue := values[0]
	maxValue := values[0]
	for _, val := range values {
		if val < minValue {
			minValue = val
		}
		if val > maxValue {
			maxValue = val
		}
	}
	return minValue, maxValue
}

func main() {
	var arraySize int
	fmt.Print("Please enter the size of the array: ")
	fmt.Scan(&arraySize)

	numbers := make([]int, arraySize)
	fmt.Println("Input the elements:")
	for i := 0; i < arraySize; i++ {
		fmt.Scan(&numbers[i])
	}

	minValue, maxValue := getMinMax(numbers)
	sort.Ints(numbers)

	fmt.Printf("Minimum value = %d\n", minValue)
	fmt.Printf("Maximum value = %d\n", maxValue)
	fmt.Print("Sorted array = ")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
