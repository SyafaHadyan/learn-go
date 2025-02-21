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

func averageNumber(num ...int) float32 {
	result := float32(sumNumber(num...)) / float32(len(num))

	return result
}

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sumNumber(number...))
	// fmt.Println(averageNumber(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(averageNumber(number...))
}
