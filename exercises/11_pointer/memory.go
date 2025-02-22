package main

import (
	"fmt"
)

func changeValue(input *int, target int) {
	*input = target
}

func main() {
	firstData := 2
	secondData := &firstData
	fmt.Printf("firstData %d\nsecondData %v\n", firstData, secondData)
	fmt.Printf("firstData %d\nsecondData %d\n", firstData, *secondData)
	changeValue(&firstData, 12)
	fmt.Printf("\nfirstData change\n\n")
	fmt.Printf("firstData %d\nsecondData %v\n", firstData, secondData)
	fmt.Printf("firstData %d\nsecondData %d\n", firstData, *secondData)
}
