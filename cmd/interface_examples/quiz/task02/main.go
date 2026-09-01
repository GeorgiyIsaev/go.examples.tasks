package main

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Meow() string {
	return "Мяу!"
}

func main() {
	var a Animal
	c := Cat{}
	a = c // ошибка компиляции
	_ = a
}

// Нужно добавить метод
func (c Cat) Speak() string {
	return "Speak!"
}
