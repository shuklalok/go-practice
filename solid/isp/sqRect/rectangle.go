package main

type rectangle struct {
	height float64
	width  float64
}

func (r *rectangle) setHight(h float64) {
	r.height = h
}
func (r *rectangle) setWidth(w float64) {
	r.width = w
}
func (r *rectangle) getArea() float64 {
	return r.height * r.width
}
func (r *rectangle) getPerimeter() float64 {
	return 2 * (r.height + r.width)
}
