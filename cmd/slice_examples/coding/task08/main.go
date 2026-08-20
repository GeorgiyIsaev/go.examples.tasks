package main

import "fmt"

//Напишите функцию MergeSorted(a, b []int) []int,
//которая принимает два отсортированных по возрастанию слайса
//и возвращает новый отсортированный слайс,
//содержащий все элементы из a и b.
//Используйте только операции срезов и append (без встроенного copy).
//Учитывайте, что исходные слайсы могут быть очень большими,
//поэтому избегайте лишних копирований.

func MergeSorted(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	if i < len(a) {
		result = append(result, a[i:]...)
	}
	if j < len(b) {
		result = append(result, b[j:]...)
	}

	return result
}

func main() {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8, 10}
	fmt.Println(MergeSorted(a, b))
}
