package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}
type Rectangle struct {
	Length, Width float64
}
type Circle struct {
	Radius float64
}
type Square struct {
	Side float64
}
type Triangle struct {
	Side1 float64
	Side2 float64
	Side3 float64
}

// func for retangle
func (rectangle Rectangle) Area() float64 {
	return rectangle.Length * rectangle.Width
}
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Length + rectangle.Width)
}
func (rectangle Rectangle) Name() string {
	return "Rectangle"
}

// func for circle
func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}
func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}
func (circle Circle) Name() string {
	return "Circle"
}

// func for square
func (square Square) Area() float64 {
	return math.Pow(square.Side, 2)
}
func (square Square) Perimeter() float64 {
	return 4 * square.Side
}
func (square Square) Name() string {
	return "Square"
}

// func for triangle
func (triangle Triangle) Area() float64 {
	// bruh
	var halfP = triangle.Perimeter() / 2
	return math.Sqrt(halfP *
		(halfP - triangle.Side1) *
		(halfP - triangle.Side2) *
		(halfP - triangle.Side3))
}
func (triangle Triangle) Perimeter() float64 {
	return triangle.Side1 + triangle.Side2 + triangle.Side3
}
func (triangle Triangle) Name() string {
	return "Triangle"
}

func ShapesCLI() {
	shapes := []Shape{
		Rectangle{Length: 5, Width: 3},
		Circle{Radius: 4},
		Square{Side: 4},
		Triangle{Side1: 3, Side2: 4, Side3: 5},
	}
	for _, s := range shapes {
		fmt.Println("Shape: ", s.Name())
		fmt.Println("Area of: ", s.Area())
		fmt.Println("Perimeter", s.Perimeter())
		fmt.Println("============")

	}
}
