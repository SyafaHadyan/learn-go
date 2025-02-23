package main

import (
	"fmt"
	"math"
)

type shapeProperty interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	side float64
}

func (a circle) area() float64 {
	return math.Pi * math.Pow(a.radius, 2)
}

func (p circle) perimeter() float64 {
	return 2 * math.Pi * p.radius
}

func (a rectangle) area() float64 {
	return math.Pow(a.side, 2)
}

func (p rectangle) perimeter() float64 {
	return 2 * (p.side + p.side)
}

func main() {
	circleShape := circle{14}
	fmt.Printf("Area\t\t: %.3f\nPerimeter\t: %.3f\n", circle.area(circleShape), circle.perimeter(circleShape))

	rectangleShape := rectangle{12}
	fmt.Printf("Area\t\t: %.3f\nPerimeter\t: %.3f\n", rectangle.area(rectangleShape), rectangle.perimeter(rectangleShape))
}
