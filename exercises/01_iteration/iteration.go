package main

import "fmt"

func Repeat(input string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result = result + input
	}
	return result
}

func main() {
	fmt.Println(Repeat("a", 5))
}
