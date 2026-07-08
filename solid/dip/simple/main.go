package main

import "fmt"

//--------------------------------
type english struct{}

func (e english) greet() string {
	return "Hello!"
}

//--------------------------------
type french struct{}

func (f french) greet() string {
	return "Bonjour!"
}

//--------------------------------
// High-level module depends directly on concrete types.
// Adding a new language requires a new function.
// This violates DIP - high-level modules should not depend on concrete types
func greetEnglish(e english) {
	fmt.Println(e.greet())
}

func greetFrench(f french) {
	fmt.Println(f.greet())
}

//--------------------------------
func main() {
	e := english{}
	greetEnglish(e)
	
	f := french{}
	greetFrench(f)
}