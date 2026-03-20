package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Hight float64
	Width float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Hight)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Hight * rectangle.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
