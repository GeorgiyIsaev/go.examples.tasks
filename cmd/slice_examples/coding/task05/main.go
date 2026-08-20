package main

import "fmt"

//Задача 5 (Средняя)
//Напишите функцию Rotate(s []int, k int),
//которая выполняет циклический сдвиг
//элементов слайса s вправо на k позиций.
//Изменения должны производиться на месте
//(без выделения нового массива).
//Обработайте случаи, когда k больше длины слайса
//(используйте остаток от деления).

func Rotate(s []int, k int) {
	if len(s) == 0 {
		return
	}
	k %= len(s)
	if k == 0 {
		return
	}

	reverse(s) //весь слайс
	fmt.Println("весь слайс", s)
	reverse(s[:k]) //перед к
	fmt.Println("перед k", s)
	reverse(s[k:]) //после л
	fmt.Println("после k", s)

}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Оригинал", s)

	Rotate(s, 2)
	fmt.Println("Результат", s)
}
