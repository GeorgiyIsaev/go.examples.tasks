package main

import "fmt"

//Задача 1. Интерфейс fmt.Stringer для пользовательского типа
//Условие:
//Создайте тип Person с полями Name (string) и Age (int).
//Реализуйте интерфейс fmt.Stringer так,
//чтобы метод String() возвращал строку в формате "Имя (Возраст лет)"
//(например, "Алексей (30 лет)").
//Напишите функцию PrintStringer(s fmt.Stringer),
//которая принимает любой Stringer и выводит его строковое представление в stdout.
//Продемонстрируйте работу.

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func PrintStringer(s fmt.Stringer) {
	fmt.Println(s.String())
	fmt.Println(s)
}

func main() {
	p := &Person{"Bob", 20}

	fmt.Println(p)
	PrintStringer(p)
}
