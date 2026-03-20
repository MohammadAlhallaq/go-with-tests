package structs

import "math"

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

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
