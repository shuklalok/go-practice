package main

import "fmt"

type shape interface {
	area() float64
}
type threeD interface {
	shape
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

func areaSum(shapes ...shape) (sum float64) {
	for _, s := range shapes {
		sum += s.area()
	}
	return
}

func volumeSum(shapes ...threeD) (sum float64) {
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return
}

func main() {
	c := cube{side: 5}
	s := square{side: 5}

	fmt.Println(areaSum(c, s))
	fmt.Println(volumeSum(c))
}
