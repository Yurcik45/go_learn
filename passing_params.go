package main

import (
    "fmt"
)

// Shape defines the Shape interface with a Disappear method.
type Shape interface {
    Disappear()
}

// Circle implements the Shape interface.
type Circle struct {
    Radius float64
}

// Disappear sets the radius of a Circle to 0.
func (c *Circle) Disappear() {
	c.Radius = 0
}

func Destruct(s Shape) {
    s.Disappear()
}

func main() {
    circle := Circle{Radius: 3.0} // Use a pointer to Circle
	Destruct(&circle)
    fmt.Printf("Circle Radius after destruction: %.2f\n", circle.Radius)
}
