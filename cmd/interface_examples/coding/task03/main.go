package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

//Type Assertion
//Напишите функцию, которая принимает интерфейс interface{}
//и печатает его тип с помощью type switch.
//Если это int, выведите квадрат числа; если string,
//выведите длину строки; если bool,
//выведите противоположное значение;
//для других типов выведите "Неизвестный тип".

func SwitchInterface(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Println(v * v)
	case string:
		fmt.Println(utf8.RuneCountInString(v))
	case bool:
		fmt.Println(!v)
	default:
		fmt.Println("Неизвестный тип", reflect.TypeOf(v))
	}
}

func main() {
	SwitchInterface(10)
	SwitchInterface("Три")
	SwitchInterface(true)
	SwitchInterface(0.0)
	SwitchInterface(nil)
}
