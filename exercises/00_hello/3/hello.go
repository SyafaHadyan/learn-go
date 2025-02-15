package main

import "fmt"

const helloWord string = "Hello, "

func Hello(name string) string {
  if name == "" {
      name = "World"
  }
	return helloWord + name
}

func main() {
	fmt.Println(Hello("Person"))
}
