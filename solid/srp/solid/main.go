package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius   float64
	typeName string
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length   float64
	typeName string
}

func (s square) area() float64 {
	return s.length * s.length
}

func printAreaFloat64(a float64, object string) {
	fmt.Printf("The area of %s is :%f\n", object, a)
}

func printAreaInt(a int, object string) {
	fmt.Printf("The area of %s is :%d\n", object, a)
}

func main() {
	c := circle{
		radius:   5,
		typeName: "circle",
	}
	s := square{
		length:   5,
		typeName: "square",
	}
	printAreaFloat64(c.area(), c.typeName)
	printAreaInt(int(s.area()), s.typeName)
}
