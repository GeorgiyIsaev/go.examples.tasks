package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() {

	//d = &Dog{} // но можно заменить но в интефейсе будет также nil
	//сохраняется только содержимое, так как адрес новый изменение не проихойдет
	//_ = *d     // будет паника если обратимся к nil
	fmt.Println("Гав")
	//метод не обращается к полям структуры паники не будет
}

func Describe(a Animal) {
	if a == nil {
		fmt.Println("nil интерфейс")
	} else {
		fmt.Println("не nil интерфейс")
		a.Speak() // если a хранит nil и Speak() обращается к указателю - паника
		//Speak() не обращается успешно вызовит Гав при nil
	}
}

func main() {
	var d *Dog // nil
	Describe(d)
}
