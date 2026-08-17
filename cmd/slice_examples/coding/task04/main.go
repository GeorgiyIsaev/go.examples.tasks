package main

import "fmt"

//Задача 4 (Средняя)
//Напишите функцию Unique(s []int) []int,
//которая возвращает новый слайс,
//содержащий только уникальные элементы из s,
//сохраняя порядок первого вхождения.
//Например: [1,2,2,3,1] → [1,2,3].
//Оптимизируйте алгоритм (можно использовать map для проверки наличия).

func Unique(s []int) []int {
	result := make([]int, 0, len(s))

	for _, v := range s {
		found := false
		for _, v2 := range result {
			if v == v2 {
				found = true
				break
			}
		}

		if !found {
			result = append(result, v)
		}
	}
	return result
}
func UniqueWMap(s []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0, len(s))

	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 2, 3, 1}
	b := []int{6, 5, 6, 5, 3}
	c := Unique(a)
	d := Unique(b)
	cm := UniqueWMap(a)
	dm := UniqueWMap(b)

	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("cm=", cm)
	fmt.Println("dm=", dm)
}
