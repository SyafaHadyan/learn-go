package main

import "fmt"

const helloWordEnglish string = "Hello, "
const helloWordSpanish string = "Hola, "
const helloWordFrench string = "bonjour, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    switch language {
    case "Spanish":
        return helloWordSpanish + name
    case "French":
        return helloWordFrench + name
    default:
        return helloWordEnglish + name
    }
}

func main() {
	fmt.Println(Hello("Person", "English"))
}
