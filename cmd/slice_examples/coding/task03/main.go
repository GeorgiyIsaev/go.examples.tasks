package main

import "fmt"

//Напишите функцию Concat(a, b []int) []int,
//которая объединяет два слайса целых чисел в один
//и возвращает новый слайс.

func Concat(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	result = append(result, a...)
	result = append(result, b...)

	return result
}

func main() {
	a := []int{11, 12, 13, 14, 15}
	b := []int{331, 312, 313, 314, 315}
	c := Concat(a, b)

	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
}
