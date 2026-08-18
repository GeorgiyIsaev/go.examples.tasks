package main

import "fmt"

func main() {
	s := "Привет"
	fmt.Println(len(s)) //12 количество байт не рун
	fmt.Println(s[0])   //вернут первый байт часть руны, какое-то число
}
