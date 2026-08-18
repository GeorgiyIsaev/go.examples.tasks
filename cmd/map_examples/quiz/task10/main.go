package main

import "fmt"
import "reflect"

func main() {
	m1 := map[int]string{1: "a", 2: "b"}
	m2 := map[int]string{2: "b", 1: "a"}
	//fmt.Println(m1 == m2)                  // строка А
	// невыполнимо (мап не сравниваемый тип)
	fmt.Println(reflect.DeepEqual(m1, m2)) // строка Б
	// поэлементное сравнение true
}
