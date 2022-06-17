package main

import (
	"fmt"
)

type car struct {
	color string
	size  int
}

func newCar() car {
	return car{color: "Opaque", size: 1}
}

func (c *car) paint(color string) {
	fmt.Printf("Painting the car with %s color\n", color)
	c.color = color
}

func (c *car) resize(size int) {
	fmt.Printf("Resizing the car with %d size\n", size)
	c.size = size
}

func main() {
	c := newCar()
	fmt.Printf("The new car is %s in color and it's size is %d\n", c.color, c.size)
	c.paint("White")
	fmt.Printf("The new car is %s in color and it's size is %d\n", c.color, c.size)
	c.resize(2)
	fmt.Printf("The new car is %s in color and it's size is %d\n", c.color, c.size)
}
