package main

import "fmt"

type Animal interface {
	Laugh() string
	Cry() string
	Tails() int
}

type Human struct {
	eyes  int
	hands int
	tail  int
}

type Monkey struct {
	eyes  int
	hands int
	tail  int
}

func NewHuman() *Human {
	return &Human{eyes: 2, hands: 2, tail: 0}
}

func (h Human) Laugh() string {
	return "Hahaha"
}

func (h Human) Cry() string {
	return "Buhuhu"
}

func (h Human) Tails() int {
	return h.tail
}

func NewMonkey() *Monkey {
	return &Monkey{eyes: 2, hands: 2, tail: 1}
}

func (m Monkey) Laugh() string {
	return "Hohoho"
}

func (m Monkey) Cry() string {
	return "Pipipi"
}

func (m Monkey) Tails() int {
	return m.tail
}

func main() {
	var a Animal

	h := NewHuman()
	fmt.Println("Human has", h.eyes, h.hands, h.tail)
	fmt.Println("Human is laughing", h.Laugh())
	fmt.Println("Human is crying", h.Cry())
	a = h
	fmt.Println("Animal has tails", a.Tails())
	m := NewMonkey()
	a = m
	fmt.Println("Monkey has", m.eyes, m.hands, m.tail)
	fmt.Println("Monkey is laughing", a.Laugh())
	fmt.Println("Monkey is crying", a.Cry())
	fmt.Println("Animal has tails", m.Tails())

}
