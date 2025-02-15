package main

import (
	"fmt"
	"strings"
)

func Repeat(input string, count int) string {
	var result strings.Builder
	for i := 0; i < count; i++ {
		result.WriteString(input)
	}
	return result.String()
}

func main() {
	fmt.Println(Repeat("a", 5))
}
