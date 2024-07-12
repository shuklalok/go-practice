package main

import "fmt"

type squareType interface {
	setHeight(float64)
	getArea() float64
}

type square struct {
	height float64
}

func (s *square) setHeight(h float64) {
	s.height = h
}
func (s *square) getArea() float64 {
	return s.height * s.height
}

type rectangle struct {
	square
	width float64
}

func (r *rectangle) setHeight(h float64) {
	r.square.height = h
}

func (r *rectangle) setWidth(w float64) {
	r.width = w
}

func (r *rectangle) getArea() float64 {
	return r.square.height * r.width
}

func area(r squareType) float64 {
	return r.getArea()
}

func main() {
	r := rectangle{}
	r.setHeight(5)
	r.setWidth(4)
	fmt.Println(area(&r))
	s := square{}
	s.setHeight(3)
	fmt.Println(area(&r))
}
