package main

import "fmt"

func main() {
	m := map[[]int]string{}
	m[[]int{1, 2}] = "hello"
	//невозможно слайс не сравниваемы тип
	//нельзя использовать как ключ
	fmt.Println(m)
}
