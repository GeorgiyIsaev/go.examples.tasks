package main

import "fmt"

// Напишите функцию copyMap(src map[string][]int) map[string][]int,
// которая создаёт полную глубокую копию map.
// Затем ответьте на вопрос: что выведет следующий код в main?

func copyMap(src map[string][]int) map[string][]int {
	result := make(map[string][]int, len(src))
	for k, v := range src {
		sliceCopy := make([]int, len(v))
		copy(sliceCopy, v) // копирование по значению, не по памяти
		result[k] = sliceCopy
	}
	return result
}

func main() {
	src := map[string][]int{
		"a": {1, 2, 3},
		"b": {4, 5},
	}
	dst := copyMap(src)
	dst["a"][0] = 100
	fmt.Println(src["a"]) //изменение dst, не повлияет на src [1 2 3]
}
