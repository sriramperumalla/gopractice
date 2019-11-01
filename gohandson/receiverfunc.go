package gohandson

import (
	"fmt"
	"math"
)

// Interface with methods defined in it and implemented by struct receiver type which is assigned to variables r and c resply. Area method is set on these variables.
type Shape interface {
	Area() float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() (Area float64) {
	Area = r.width * r.height
	return
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func ImplementInterface() {
	var s Shape
	s = Rect{3.0, 4.0}
	s = Circle{4.0}
	fmt.Println("--- Reciever Function Impelmentation ---")
	fmt.Printf("Area of Rectangle (Type: %T, Value: %v, Area: %0.2f", s, s, s.Area())
	fmt.Printf("Aread of Circle (Type: %T, Value:%v, Area: %0.2f", s, s, s.Area())
	fmt.Println("--- Reciever Function Impelmentation ---")
}
