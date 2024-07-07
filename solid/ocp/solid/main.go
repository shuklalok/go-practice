package main

import (
	"fmt"
)

type shape interface {
	draw()
}

type circle struct {
	radius float64
}

func (c circle) draw() {
	fmt.Println("Drawing a circle")
}

type square struct {
	side float64
}

func (s square) draw() {
	fmt.Println("Drawing a square")
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) draw() {
	fmt.Println("Drawing a rectangle")
}

type draw struct {
}

func (d draw) drawShape(s ...shape) {
	for _, shape := range s {
		shape.draw()
	}
}

func main() {
	c := circle{radius: 5}
	s := square{side: 5}
	r := rectangle{length: 5, width: 5}
	d := draw{}
	d.drawShape(c, s, r)
}
