package main

import "fmt"

func Sum(numbers []int) int {
	var result int
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func SumAll(numbers ...[]int) []int {
	var result []int = make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			result[i] = Sum(numbers[i])
		}
	}
	return result
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(SumAll([]int{1, 2, 3}, []int{2, 3, 4}))
}
