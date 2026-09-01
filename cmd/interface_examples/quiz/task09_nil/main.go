package main

import (
	"fmt"
	"reflect"
)

type Speaker interface {
	Speak()
}

type Person struct {
	Name string
}

func (p *Person) Speak() {
	fmt.Println("Hello, I'm", p.Name)
}

func main() {
	var s Speaker

	var p *Person // p == nil
	s = p

	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
		// s.Speak() // panic: nil pointer dereference при вызове
	}
	//Значение равно nil, но сам интерфейс не nil

	//Решение 1. Если известно значение
	if p, ok := s.(*Person); ok && p != nil {
		// значение не nil, можно безопасно вызывать методы
		p.Speak()
	} else {
		fmt.Println("значение nil или тип не соответствует")
	}

	//Решение 2. Использовать рефлексию
	if s != nil {
		v := reflect.ValueOf(s)
		// Проверяем, что динамическое значение может быть nil и действительно равно nil
		if v.Kind() == reflect.Ptr && v.IsNil() {
			fmt.Println("динамическое значение nil")
		} else {
			fmt.Println("динамическое значение не nil")
			// можно безопасно вызывать методы через интерфейс
			s.Speak()
		}
	} else {
		fmt.Println("сам интерфейс nil")
	}
}
