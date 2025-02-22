package main

import "fmt"

type person struct {
	name string
	age  int
}

func (person person) getPersonInfo() (string, int) {
	return person.name, person.age
}

func main() {
	personList := []person{
		{
			"person0",
			10,
		},
		{
			"person1",
			11,
		},
		{
			"person2",
			12,
		},
		{
			"person3",
			13,
		},
	}

	for _, get := range personList {
		name, age := get.getPersonInfo()
		fmt.Printf("Name\t: %s\nAge\t: %d\n\n", name, age)
	}
}
