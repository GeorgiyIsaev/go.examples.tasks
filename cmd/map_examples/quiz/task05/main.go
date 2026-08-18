package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {"Alice", 25},
	}
	// Попытка изменить поле
	m[1].Age = 26
	//Код не скомпилируется с ошибкой:
	//cannot assign to struct field m[1].Age in map.
	//Элемент map не адресуем, поэтому напрямую изменить поле нельзя.
	//Нужно скопировать структуру, изменить и записать обратно:
	fmt.Println(m[1].Age) //26

	//ПРАВИЛЬНО
	u := m[1]
	u.Age = 26
	m[1] = u
	fmt.Println(m[1].Age) //26
}
