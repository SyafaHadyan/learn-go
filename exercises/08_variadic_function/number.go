package main

import (
	"fmt"
)

func sumNumber(num ...int) int {
	var result int
	for i := 0; i < len(num); i++ {
		result += num[i]
	}

	return result
}

func averageNumber(num ...int) float64 {
	return float64(sumNumber(num...)) / float64(len(num))
}

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Number: %v\n", number)
	fmt.Printf("Sum: %d\n", sumNumber(number...))
	fmt.Printf("Average: %.3f\n", averageNumber(number...))
}
