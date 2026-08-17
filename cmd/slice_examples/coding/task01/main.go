package main

import "fmt"

//Напишите функцию Reverse(s []int) []int,
//которая принимает слайс целых чисел
//и возвращает новый слайс с элементами в обратном порядке.
//Исходный слайс не должен изменяться.

func Reverse(s []int) []int {
	result := make([]int, 0, len(s))
	i := 0
	for i < len(s) {
		result = append(result, s[len(s)-i-1])
		i++
	}
	return result
}

func main() {
	a := []int{11, 12, 13, 14, 15}
	b := Reverse(a)

	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
