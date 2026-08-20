package main

import "fmt"

//Напишите функцию StringLengths(s []string) []int,
//которая принимает слайс строк
//и возвращает слайс целых чисел — длин каждой строки,
//но только для тех строк, длина которых больше 3.
//Если строка короче или равна 3, она пропускается
//(в результирующем слайсе её длина не появляется).
//Сохраняйте порядок.

func StringLengths(s []string) []int {
	result := []int{}
	for _, str := range s {
		if len(str) > 3 {
			result = append(result, len(str))
		}
	}
	return result
}

func main() {
	strings := []string{"hello", "go", "world", "a", "golang", "yes", "four"}
	lengths := StringLengths(strings)
	fmt.Println(lengths)
}
