package main

import (
	"fmt"
	"sort"
)

//Интерфейс sort.Interface
//Дан тип Person (поля Name string, Age int) и тип People – срез []Person.
//Реализуйте интерфейс sort.Interface для типа People так,
//чтобы сортировка происходила по возрасту (по возрастанию).
//Напишите функцию SortAndPrint(p People), которая сортирует срез с помощью sort.Sort
//и выводит результат в удобочитаемом виде (например, каждой строкой "Name - Age").
//Продемонстрируйте работу.

type Person struct {
	Name string
	Age  int
}

type People []Person

// Реализация sort.Interface
func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	people := People{
		{"Bob", 19},
		{"Alice", 22},
		{"Yl", 18},
	}
	fmt.Println(people)

	sort.Sort(people)
	fmt.Println(people)

}
