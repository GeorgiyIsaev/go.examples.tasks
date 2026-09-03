package main

import "fmt"

type MyInterface interface {
	M()
}

type T struct{}

func (t *T) M() {}

func main() {
	var i MyInterface
	var t *T // nil
	i = t
	fmt.Println(i == nil) // false
	// i.M() // паника: runtime error: invalid memory address or nil pointer dereference

	p, ok := i.(*T)       //явное приведение
	fmt.Println(p == nil) //true
	fmt.Println(ok)       //true

}
