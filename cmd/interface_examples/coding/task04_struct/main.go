package main

import "fmt"

//Встраивание структур
//Создайте структуру Animal с полями Name и Age.
//Создайте структуру Dog с встраиванием Animal
//и дополнительным полем Breed.
//Напишите метод Speak() для Dog, который выводит "Гав!".
//Создайте экземпляр Dog и вызовите метод, а также обратитесь к полям из Animal.

type Animal struct {
	Name string
	Age  int
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "Сказал Гав")
}

func main() {
	dog1 := Dog{
		Animal: Animal{Name: "Кусь", Age: 10},
		Breed:  "Лапусь",
	}
	dog1.Speak()

	// Способ 2 — порядковое перечисление
	dog2 := Dog{Animal{"Пес", 10}, "Барбос"}
	dog2.Speak()
}
