package main

import (
	"fmt"
	"reflect"
)

//Калькулятор площади с интерфейсами
//Реализуйте интерфейс Figure с методом Area() float64.
//Создайте структуры: Square (сторона), Rectangle (ширина, высота),
//Circle (радиус). Напишите функцию, которая принимает
//слайс Figure и возвращает сумму площадей.
//Добавьте функцию, которая находит фигуру с максимальной площадью.

type Figure interface {
	Area() float64
}

type Square struct {
	a float64
}

func (s Square) Area() float64 {
	return s.a * s.a
}

type Rectangle struct {
	a float64
	b float64
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}

type Circle struct {
	r float64
}

func (с Circle) Area() float64 {
	return 3.14 * с.r * с.r
}

func Sum(figures []Figure) float64 {
	var s float64
	for _, figure := range figures {
		s += figure.Area()
	}
	return s
}

func MaxFigure(figures []Figure) Figure {
	var max float64
	var f Figure
	for _, figure := range figures {
		if figure.Area() > max {
			max = figure.Area()
			f = figure
		}
	}
	return f
}

func main() {
	figures := []Figure{Square{10}, Rectangle{8, 9}, Circle{12}}
	sum := Sum(figures)
	fmt.Println("Сумма", sum)
	maxF := MaxFigure(figures)
	fmt.Println("Фигура", reflect.TypeOf(maxF), "Сумма", maxF.Area())
}
