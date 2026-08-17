package main

import "fmt"

// Что выведем
func main() {
	a := []int{1, 2, 3} // len: 3 cap:3
	b := a              // len: 3 cap: 3

	b := append(a, 100)

	fmt.Println(a) // 1 2 3
	fmt.Println(b) // 1 2 3 100
}
