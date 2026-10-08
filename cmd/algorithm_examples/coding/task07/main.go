package main

import "fmt"

//7. Слить два отсортированных среза
//Условие: Даны два отсортированных по возрастанию среза []int. Вернуть новый срез, содержащий все элементы обоих в отсортированном порядке. Не изменять входные срезы.
//Примеры:
//•	[1,3,5], [2,4,6] → [1,2,3,4,5,6]
//•	[1,2,3], [] → [1,2,3]
//•	[], [] → []
//•	[1,1,2], [1,3] → [1,1,1,2,3]
//Сигнатура:
//go
//func MergeSorted(a, b []int) []int

func MergeSorted(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	// Если один срез закончился, добавляем остаток второго.
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)

	return res
}

func main() {
	fmt.Println(MergeSorted([]int{1, 3, 5}, []int{2, 4, 6}))
	// [1 2 3 4 5 6]

	fmt.Println(MergeSorted([]int{1, 1, 2}, []int{1, 3}))
	// [1 1 1 2 3]
}
