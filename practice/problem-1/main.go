package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius

}

func printArea(shape Shape) {
	fmt.Printf("Area is: %.2f\n", shape.Area())
}

func main() {
	shapes := []Shape{
		Rectangle{width: 5, height: 3},
		Circle{radius: 5},
	}
	for _, s := range shapes {
		printArea(s)
	}
}
