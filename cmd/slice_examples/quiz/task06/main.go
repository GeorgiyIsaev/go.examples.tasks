package main

import "fmt"

// Вопрос: Что будет выведено?
func main() {
	var s []int
	if s == nil {
		fmt.Println("nil") //nil
	} else {
		fmt.Println("not nil")
	}
	fmt.Println(len(s), cap(s)) //len 0 cap 0
}
