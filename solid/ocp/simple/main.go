package main

import (
	"fmt"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type draw struct{}

func (d draw) drawShape(s ...interface{}) {
	for _, shape := range s {
		switch shape.(type) {
		case circle:
			fmt.Println("Drawing a circle")
		case square:
			fmt.Println("Drawing a square")
		}
	}
}

func main() {
	c := circle{radius: 5}
	s := square{side: 5}
	d := draw{}
	d.drawShape(c, s)
}
