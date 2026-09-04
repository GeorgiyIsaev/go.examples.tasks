package main

import "fmt"

//Упражнение: Интерфейс для транспортных средств
//Определите интерфейс Vehicle с методами Start() string и Stop() string.
//Реализуйте структуры Car и Bicycle, каждая со своим методом Start и Stop.
//Напишите функцию, которая принимает Vehicle и вызывает оба метода, выводя результаты.

type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct{}

func (c Car) Start() string {
	return "Машина поехала!"
}

func (c Car) Stop() string {
	return "Машина остановилась!"
}

type Bicycle struct{}

func (b Bicycle) Start() string {
	return "Велосипед поехал!"
}

func (b Bicycle) Stop() string {
	return "Велосипед остановился!"
}

func Road(v Vehicle) {
	fmt.Println(v.Start())
	fmt.Println(v.Stop())
}

func main() {
	var car, bicycle Vehicle
	car = Car{}
	bicycle = Bicycle{}

	Road(car)
	Road(bicycle)
}
