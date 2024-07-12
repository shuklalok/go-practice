package main

import (
	"fmt"
)

type shape interface {
	area() float64
	volume() float64
}

type cube struct {
	side float64
}

func (c cube) area() float64 {
	return 6 * (c.side * c.side)
}

func (c cube) volume() float64 {
	return c.side * c.side * c.side
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) volume() float64 {
	return 0
}

func areaSum(shape ...shape) (sum float64) {
	for _, s := range shape {
		sum += s.area()
	}
	return
}

func volumeSum(shape ...shape) (sum float64) {
	for _, s := range shape {
		sum += s.area() + s.volume()
	}
	return
}
func main() {
	c := cube{side: 5}
	s := square{side: 5}

	fmt.Println(areaSum(c, s))
	fmt.Println(volumeSum(c, s))
}
