package main

import (
    "fmt"
    "strings"
    "math/rand"
    "time"
)

func printName(greet string, name []string) {
    for i := 0; i < len(name); i++ {
        fmt.Print(greet + " ")
        fmt.Println(name[i])
    }
}

func printNameJoin(name []string) {
    fmt.Println(strings.Join(name, " "))
}

func getRandomNumber() int {
    return rand.New(rand.NewSource(time.Now().Unix())).Int()
}

func main() {
    var names = []string{"Person0", "Person1"}
    printName("testGreet", names)
    printNameJoin(names)
    fmt.Println(getRandomNumber())
}

