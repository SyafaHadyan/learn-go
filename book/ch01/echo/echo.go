package main

import (
    "fmt"
    "strings"
    "os"
)

func main() {
    fmt.Println("Command: " + os.Args[0])
    fmt.Println("Arguments: ")
    fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Printf("Argument count: %d\n", len(os.Args) - 1)
}

