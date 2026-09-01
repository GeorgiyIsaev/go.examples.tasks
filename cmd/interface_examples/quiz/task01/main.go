package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Гав!"
}

func main() {
	var a Animal
	d := Dog{}
	a = d
	fmt.Println(a.Speak())
}
