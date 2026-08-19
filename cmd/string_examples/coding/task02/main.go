package main

import (
	"fmt"
	"sort"
	"strings"
)

//Напишите функцию IsAnagram(a, b string) bool,
//которая проверяет, являются ли две строки анаграммами
//(игнорируя регистр и пробелы).

func IsAnagram(a, b string) bool {
	//единая форма
	a = strings.ToLower(strings.ReplaceAll(a, " ", ""))
	b = strings.ToLower(strings.ReplaceAll(b, " ", ""))
	if len(a) != len(b) {
		return false
	}

	//срез рун
	aRunes := []rune(a)
	bRunes := []rune(b)

	// Сортируем срезы
	sort.Slice(aRunes, func(i, j int) bool { return aRunes[i] < aRunes[j] })
	sort.Slice(bRunes, func(i, j int) bool { return bRunes[i] < bRunes[j] })

	// Сравниваем отсортированные срезы как строки
	return string(aRunes) == string(bRunes)

}
func main() {
	r1 := IsAnagram("listen", "silent") // true
	fmt.Println(r1)
	r2 := IsAnagram("A gentleman", "Elegant man") // true
	fmt.Println(r2)
}
