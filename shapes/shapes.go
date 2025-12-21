package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
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

// func for circle
func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}
func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

// func for square
func (square Square) Area() float64 {
	return math.Pow(square.Side, 2)
}
func (square Square) Perimeter() float64 {
	return 4 * square.Side
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
func (triagnle Triangle) Perimeter() float64 {
	return triagnle.Side1 + triagnle.Side2 + triagnle.Side3
}

func ShapesCLI() {
	var s Shape

	s = Rectangle{Length: 5, Width: 3}
	fmt.Println("Rectangle area: ", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())

	s = Circle{Radius: 4}
	fmt.Println("Circle area", s.Area())
	fmt.Println("Circle perimeter:", s.Perimeter())

	s = Square{Side: 4}
	fmt.Println("Square area", s.Area())
	fmt.Println("Square perimeter:", s.Perimeter())

	s = Triangle{Side1: 3, Side2: 4, Side3: 5}
	fmt.Println("Triangle area", s.Area())
	fmt.Println("Triangle perimeter:", s.Perimeter())
}
