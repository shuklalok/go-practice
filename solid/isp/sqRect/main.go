package main

import (
	"fmt"
)

func area(r rectangleType) float64 {
	return r.getArea()
}
func perimeter(r rectangleType) float64 {
	return r.getPerimeter()
}

func main() {
	r := rectangle{}
	r.setHight(5)
	r.setWidth(5)
	fmt.Println(area(&r))
	fmt.Println(perimeter(&r))

	s := square{}
	s.setHeight(5)
	s.setWidth(5)
	fmt.Println(area(&s))
	fmt.Println(perimeter(&s))
}
