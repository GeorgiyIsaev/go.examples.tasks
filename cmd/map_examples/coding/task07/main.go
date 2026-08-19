package main

import "fmt"

// Напишите функцию isSubset(sub, super map[string]int) bool,
// которая проверяет, что все ключи из sub присутствуют в super
// и значения по этим ключам совпадают.

func isSubset(sub, super map[string]int) bool {
	if len(sub) > len(super) {
		return false
	}

	result := true

	for k, v := range sub {
		val := super[k]
		if val != v {
			return false
		}
	}

	return result
}

func main() {
	sub := map[string]int{"a": 1, "b": 2}
	super := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(isSubset(sub, super)) // true

	sub2 := map[string]int{"a": 1, "b": 5}
	fmt.Println(isSubset(sub2, super)) // false (значение не совпадает)
}
