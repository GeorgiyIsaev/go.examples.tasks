package main

import "fmt"

func main() {
	a := "hello"
	b := "hello"
	c := "world"
	fmt.Println(a == b) //true - сравнение по символьное по значению
	fmt.Println(a == c) //false
}
