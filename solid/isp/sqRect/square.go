package main

type square struct {
	rectangle
}

func (s square) setHeight(h float64) {
	s.height = h
	s.width = h
}
func (s square) setWidth(w float64) {
	s.width = w
	s.height = w
}

func (s square) getArea() float64 {
	return s.height * s.width
}
func (s square) getPerimeter() float64 {
	return 4 * s.height
}
