package main

import "fmt"

// Что выведем:
func main() {
	a := []int{1, 2, 3} // len: 3 cap:3 //При создании cap равен len, заранее ничего не выделяем
	b := a              // len: 3 cap: 3

	b[0] = 100

	fmt.Println(a) // 100 2 3
	fmt.Println(b) // 100 2 3
}
