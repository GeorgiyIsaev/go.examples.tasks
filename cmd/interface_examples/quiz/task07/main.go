package main

import "fmt"

type I interface{}

type T struct{}

func main() {
	var a, b I
	a = 42
	b = 42
	fmt.Println(a == b) // true

	a = T{}
	b = T{}
	fmt.Println(a == b) // true
	// a == nil, b == nil - true
	fmt.Println(a)
	fmt.Println(b)
	//сравнение по значениям внутри структур

	a = &T{}
	b = &T{}
	fmt.Println(a == b) // false (разные адреса)
	fmt.Println(a)
	fmt.Println(b)
	//сравнения по месту хранения

	// a == nil, b == nil - true
}
