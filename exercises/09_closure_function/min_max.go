package main

import (
	"fmt"
)

func main() {
	getMinMax := func(numberInput []int) (int, int) {
		var min int
		var max int
		for i := 0; i < len(numberInput); i++ {
			if numberInput[i] > max {
				max = numberInput[i]
			} else if numberInput[i] < min {
				min = numberInput[i]
			}
		}
		return min, max
	}
	number := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	min, max := getMinMax(number)
	fmt.Printf("Number\t: %v\nMin\t: %d\nMax\t: %d\n", number, min, max)
}
