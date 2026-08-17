package main

import "fmt"

// Что произойдёт? Вызовет ли панику? Объясните.
func main() {
	a := []int{1, 2, 3} //len3 cap3
	b := a[1:2]         //[2] //len1 cap2
	b = b[:3]           //[2,3,x] //len3 cap3 //выход за пределы ёмкости
	fmt.Println(b)      //паника
}
