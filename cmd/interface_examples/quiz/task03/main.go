package main

import "fmt"

func main() {
	var i interface{} = 42
	v, ok := i.(int)
	fmt.Println(v, ok) //42 true

	v2, ok := i.(string)
	fmt.Println(v2, ok) //  false

	// Попытка без проверки ok (паника)
	// v3 := i.(string) // panic: interface conversion: interface {} is int, not string
}
