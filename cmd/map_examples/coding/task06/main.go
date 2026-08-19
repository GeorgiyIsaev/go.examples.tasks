package main

import "fmt"

//Напишите функцию mergeMaps(m1, m2 map[string]int) map[string]int,
//которая объединяет две map.
//Если ключ присутствует в обеих, приоритет имеет значение из m2.
//Функция не должна изменять исходные map.

func mergeMaps(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v //перезапишем ключ
	}
	return result
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	result := mergeMaps(m1, m2) // {"a":1, "b":3, "c":4}
	fmt.Println(result)
}
