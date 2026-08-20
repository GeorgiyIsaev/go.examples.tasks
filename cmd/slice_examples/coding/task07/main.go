package main

import "fmt"

//Реализуйте функцию Chunk(s []int, size int) [][]int,
//которая разбивает слайс s на слайсы размера size
//(последний может быть меньше).
//Если size <= 0 или s пуст, верните пустой слайс слайсов.
//Постарайтесь минимизировать выделения памяти
//(используйте только один дополнительный слайс для результата).

func Chunk(s []int, size int) [][]int {
	if size <= 0 || len(s) == 0 {
		return [][]int{}
	}

	n := len(s)
	chunks := (n + size - 1) / size
	result := make([][]int, 0, chunks)

	for i := 0; i < n; i += size {
		end := i + size
		if end > n {
			end = n
		}
		result = append(result, s[i:end])
	}
	return result

}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Chunk(s, 2))
}
