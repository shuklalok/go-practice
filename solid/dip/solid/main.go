package main

type Greeter interface {
	Greet() string
}

type EnglishGreeter struct{}

func (e EnglishGreeter) Greet() string {
	// English greeting details
	return "Hello!"
}

type frenchGreeter struct{}

func (f frenchGreeter) Greet() string {
	// French greeting details
	return "Bonjour!"
}

func greet(greeter Greeter) {
	// Abstraction
	println(greeter.Greet())
}

func main() {
	e := EnglishGreeter{}
	greet(e)
	f := frenchGreeter{}
	greet(f)
}
