package main

import (
	"fmt"
)

type person interface {
	getName() string
}

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	salary float64
}

func (t teacher) getName() string {
	return t.name
}

type student struct {
	human
	grade int
}

func (s student) getName() string {
	return s.name
}

type printer struct{}

func (p printer) namePrinter(person person) {
	fmt.Println("Name: " + person.getName())
}

func main() {
	h := human{name: "Human's Name"}
	t := teacher{human: human{name: "Teacher's Name"}, salary: 10000}
	s := student{human: human{name: "Student's Name"}, grade: 10}
	p := printer{}
	p.namePrinter(h)
	p.namePrinter(t)
	p.namePrinter(s)
}
