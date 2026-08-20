package main

import "fmt"

//Реализуйте функцию AppendInt(s []int, elems ...int) []int,
//которая ведёт себя как встроенный append для слайса целых чисел.
//Она должна возвращать новый слайс, который либо использует ту же память
//(если ёмкости достаточно),
//либо выделяет новый массив с увеличенной ёмкостью
//(например, удваивает, если не хватает).
//Не используйте встроенный append в вашей реализации.
//Напишите код, который правильно управляет длиной и ёмкостью.

func AppendInt(s []int, elems ...int) []int {
	n := len(s)
	m := len(elems)

	//места достаточно
	if n+m <= cap(s) {
		s = s[:n+m] // меняем длину, и заполняем
		for i := 0; i < m; i++ {
			s[n+i] = elems[i]
		}
		return s
	}
	//удваиваем кап пока не будет больше лен
	newCap := cap(s)
	for newCap < n+m {
		newCap *= 2
	}

	//заполняем новый слайс
	result := make([]int, n+m, newCap)
	for i := 0; i < n; i++ {
		result[i] = s[i]
	}
	for i := 0; i < m; i++ {
		result[n+i] = elems[i]
	}
	return result
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println("До:", a, "len=", len(a), "cap=", cap(a))
	a = AppendInt(a, 4, 5)
	fmt.Println("После:", a, "len=", len(a), "cap=", cap(a))
}
