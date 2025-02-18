package main

import (
    "fmt"
)

func getAreaPerimeterRectanngle(sideA float64, sideB float64) (float64, float64) {
    return sideA * sideB, (2 * (sideA + sideB))
}

func main() {
    var area, perimeter = getAreaPerimeterRectanngle(23, 84.2)
    fmt.Println("Area", area, "\nPerimeter", perimeter)
}

