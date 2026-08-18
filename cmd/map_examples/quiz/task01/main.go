package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	v1 := m["a"]        //1
	v2 := m["c"]        //0 значение по умолчанию
	fmt.Println(v1, v2) //1,0
}
