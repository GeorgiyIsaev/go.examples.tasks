package main

import (
	"fmt"
	"sort"
)

// Что выведет этот код? И главное — почему?
// Напишите, как это исправить, если нужно предсказуемое поведение.
func main() {
	m := map[string]int{"A": 1, "B": 2, "C": 3}

	//Порядок итерации по map рандомизирован
	for k, v := range m {
		if v%2 == 0 {
			delete(m, k) //удаление безопасно
		}
		if k == "A" {
			m["D"] = 4 //добавление недетерминированное не безопасно
		}
	}
	fmt.Println(m)
	fmt.Println("correct")
	correct()
}

func correct() {
	m := map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println("after(m)", m)
	// Собрать все ключи в слайс, отсортировать, итерироваться по слайсу
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := m[k]
		if v%2 == 0 {
			delete(m, k)
		}
		if k == "A" {
			m["D"] = 4
		}
	}

	//newM :=

	fmt.Println("before(m)", m)
}
