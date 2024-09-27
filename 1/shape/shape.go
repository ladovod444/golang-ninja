package shape

import (
	"math"
)

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}


type Square struct {
	sideLenght float32
}

func NewSquare(length float32) Square {
	return Square{
		sideLenght: length,
	}
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}
func (s Square) Perimeter() float32 {
	return s.sideLenght * 4;
}


type Circle struct {
	radius float32
}

func NewCircle(radius float32) Circle {
	return Circle{
		radius: radius,
	}
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius;
}
