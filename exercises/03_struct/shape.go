package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    width float64
    length float64
}

type Circle struct {
    radius float64
}

type Triangle struct {
    base float64
    height float64
}

func (rectangle Rectangle) Area() float64 {
    return rectangle.width * rectangle.length
}

func (circle Circle) Area() float64 {
    return math.Pi * math.Pow(circle.radius, 2)
}

func (rectangle Rectangle) Perimeter() float64 {
    return 2 * (rectangle.width + rectangle.length)
}

func (circle Circle) Perimeter() float64 {
    return 2 * math.Pi * circle.radius
}

func (triangle Triangle) Area() float64 {
    return (triangle.base * triangle.height) / 2
}

func (triangle Triangle) Perimeter() float64 {
    return triangle.base + triangle.height + (math.Sqrt(math.Pow(triangle.base, 2) + math.Pow(triangle.height, 2)))
}

func main() {
    fmt.Println(Shape(Rectangle{width: 10, length: 12}).Area())
    fmt.Println(Shape(Rectangle{width: 12, length: 15}).Perimeter())
    fmt.Println(Shape(Circle{radius: 7}).Area())
    fmt.Println(Shape(Circle{radius: 14}).Perimeter())
    fmt.Println(Shape(Triangle{base: 12, height: 5}).Area())
    fmt.Println(Shape(Triangle{base: 3, height: 4}).Perimeter())
}
