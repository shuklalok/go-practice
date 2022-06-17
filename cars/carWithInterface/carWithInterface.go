package main

import (
	"fmt"
)

var Hatchback Cartype

type Cartype interface {
	paint(string)
	resize(int)
	getColor() string // Struct filed accessor
	getSize() int     // Struct field accessor
}

type car struct {
	color string
	size  int
}

func newCar() Cartype {
	Hatchback = &car{color: "Opaque", size: 1}
	return Hatchback
}

func (c *car) paint(color string) {
	fmt.Printf("Painting the car with %s color\n", color)
	c.color = color
}

func (c *car) resize(size int) {
	fmt.Printf("Resizing the car with %d size\n", size)
	c.size = size
}

func (c *car) getColor() string {
	return c.color
}
func (c *car) getSize() int {
	return c.size
}

func main() {
	c := newCar()
	fmt.Printf("The new car is %s in color and it's size is %d\n", c.getColor(), c.getSize())
	c.paint("White")
	fmt.Printf("The new car is %s in color and it's size is %d\n", c.getColor(), c.getSize())
	c.resize(2)
	fmt.Printf("The new car is %s in color and it's size is %d\n", c.getColor(), c.getSize())
}
