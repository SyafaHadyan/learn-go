package main

import (
	"fmt"
	"learn-go/exercises/13_package/library"
)

func main() {
	person1 := library.Person{"personName", 10}
	fmt.Printf("Name\t: %s\nAge\t: %d\n", person1.Name, person1.Age)
}
