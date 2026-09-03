package main

import "fmt"

type A interface {
	Foo()
}

type B interface {
	Bar()
}

type S struct{}

func (S) Foo() {}

func main() {
	var a A = S{}
	// b := a.(B) // panic: interface conversion: main.S is not main.B: missing method Bar
	b, ok := a.(B)
	fmt.Println(b, ok) // <nil> false //безопасный способ без паники
	//b не присвоено значение так как интерфейс не подошел
}
