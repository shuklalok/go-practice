package main

import (
	"fmt"
)

type circle struct {
	radius float64
}

// area calculates the area of a circle.
//
// It takes no parameters.
// It returns a float64 representing the area of the circle.
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type square struct {
	length float64
}

// area calculates the area of a square.
//
// It takes no parameters.
// It returns a float64 representing the area of the square.
func (s square) area() float64 {
	return s.length * s.length
}

// main is the entry point of the program.
//
// It initializes a circle with a radius of 5 and a square with a length of 5.
// Then it prints the area of the circle and the square.
func main() {
	c := circle{radius: 5}
	s := square{length: 5}
	fmt.Println(c.area())
	fmt.Println(s.area())
}
