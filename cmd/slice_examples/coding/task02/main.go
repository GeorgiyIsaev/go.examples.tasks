package main

import "fmt"

//Напишите функцию RemoveAtIndex(s []int, idx int) []int,
//которая удаляет элемент с индексом idx из слайса s
//и возвращает новый слайс с сохранением порядка остальных элементов.
//Если idx некорректен, верните исходный слайс (или пустой, по вашему усмотрению).
//Реализуйте без использования дополнительного слайса (только операции среза и append).

func RemoveAtIndex(s []int, idx int) []int {
	if idx < 0 || idx >= len(s) {
		return s
	}
	s = append(s[:idx], s[idx+1:]...)
	return s
}

func main() {
	a := []int{11, 12, 13, 14, 15}
	b := RemoveAtIndex(a, 2)

	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
