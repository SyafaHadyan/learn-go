package main

import (
	"fmt"
)

func tab(input string) {
	fmt.Printf("\t%s\n", input)
}

func spaceBefore(input string) {
	fmt.Printf("  %s\n", input)
}

func formatPrint(function func(string), input string) {
	function(input)
}

func main() {
	name := []string{"first", "second", "third", "fourth"}
	for _, input := range name {
		formatPrint(tab, input)
		formatPrint(spaceBefore, input)
	}
}
