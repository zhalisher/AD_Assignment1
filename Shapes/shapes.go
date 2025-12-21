package Shapes

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
}
type Circle struct {
	radius float64
}
type Square struct {
}
type Triangle struct {
}

func ShapesCLI() {

}
