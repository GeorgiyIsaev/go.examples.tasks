package main

import "fmt"

// Что выведем
func main() {
	a := []int{1, 2, 3} //1,2,3  len3 cap3
	b := a
	b = append(b, 4) // 1,2,3,4  len4 cap6 //пересоздание
	a[1] = 10
	fmt.Println(a) // 1, 10, 3
	fmt.Println(b) // 1, 2, 3, 4
}
