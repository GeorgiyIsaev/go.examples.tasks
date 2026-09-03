package main

import "fmt"

func main() {
	var a interface{} = 10
	var b interface{} = int32(10)
	fmt.Println(a == b) // false сравнение двух разных типов
	//даже если значение одинаковое

	var aI, bI I
	aI = A{a: 10}
	bI = B{a: 10}
	fmt.Println("aI == bI (интерфейсы):", aI == bI) // false
	aI.Quack()                                      // A: 10
	bI.Quack()                                      // B: 10
}

type A struct {
	a int
}
type B struct {
	a int
}

type I interface {
	Quack()
}

func (a A) Quack() {
	fmt.Println(a.a)
}
func (a B) Quack() {
	fmt.Println(a.a)
}
