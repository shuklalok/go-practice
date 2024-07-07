package main

import "fmt"

type livingBeings interface {
	canFly()
	laysEggs()
}

type bird struct{}

func (b bird) canFly() {
	fmt.Println("... flies")
}
func (b bird) laysEggs() {
	fmt.Println("... lays eggs")
}

type mosquito struct{}

func (m mosquito) canFly() {
	fmt.Println("... can fly")
}
func (m mosquito) laysEggs() {
	fmt.Println("... can lay eggs")
}

func itCan(l livingBeings) {
	l.canFly()
	l.laysEggs()
}

func main() {
	b := bird{}
	m := mosquito{}

	itCan(b)
	itCan(m)

	fmt.Println("But wrong!")
}
