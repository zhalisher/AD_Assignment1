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
}
type Triangle struct {
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}
func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}
func ShapesCLI() {
	var s Shape
	s = Circle{Radius: 4}
	fmt.Println("Circle Area", s.Area())
	fmt.Println("C Perimeter:", s.Perimeter())

}
