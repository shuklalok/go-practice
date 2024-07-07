package main

import "fmt"

type person interface {
	getName() string
}

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type printer struct{}

func (printer) namePrinter(person person) {
	fmt.Println("Name: " + person.getName())
}

func main() {
	h := human{name: "Alok"}
	p := printer{}
	p.namePrinter(h)
}
