package main

import "fmt"

type Rectangle struct {
	width float64
	length float64
}

func RectangleArea(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.length
}

func RectanglePerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.length)
}

func main() {
	fmt.Println(RectangleArea(Rectangle{width: 10, length: 12}))
	fmt.Println(RectanglePerimeter(Rectangle{width: 12, length: 15}))
}
